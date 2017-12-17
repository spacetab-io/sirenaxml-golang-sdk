package sirena

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sirena-agent-go/config"
	"sirena-agent-go/logger"
	"sirena-agent-go/random"
	"time"

	"github.com/tmconsulting/sirenaxml-golang-sdk/crypt"
	"github.com/tmconsulting/sirenaxml-golang-sdk/des"

	"github.com/davecgh/go-spew/spew"
)

// Client is a sirena client
type Client struct {
	net.Conn
	Key []byte
}

type NewClientOptions struct {
	Test bool // skip creating and signing symmetric key
}

// NewClient connects to Sirena (if not yet) and returns sirena client singleton
func NewClient(options ...NewClientOptions) *Client {
	config := config.Get()
	conn, err := net.Dial("tcp", config.GetSirenaAddr())
	if err != nil {
		log.Fatal(err)
	}
	client := &Client{
		Conn: conn,
		Key:  nil,
	}
	if len(options) == 0 || !options[0].Test {
		// Create symmetric key
		if err := client.CreateAndSignKey(); err != nil {
			log.Fatal(err)
		}
		// Update key every 1 hour
		go func() {
			for _ = range time.Tick(time.Hour) {
				if err := client.CreateAndSignKey(); err != nil {
					log.Fatal(err)
				}
			}
		}()
	}
	return client
}

// Request is a Sirena request
type Request struct {
	Header           *Header
	SubHeader        []byte
	Message          []byte
	MessageSignature []byte
}

// Response is a Sirena response
type Response struct {
	Header  *Header
	Message []byte
}

// CreateAndSignKey creates new DES key and signs it with Sirena
func (client *Client) CreateAndSignKey() error {
	logger := logger.Get()
	logger.Debug("CreateAndSignKey")
	// Create key as a random string of 8 characters
	var key = []byte(random.String(8))
	logger.Debugf("Trying to sign DES key %s with Sirena", key)
	// Get server public key
	config := config.Get()
	serverPublicKey, err := config.GetKeyFile(config.ServerPublicKey)
	if err != nil {
		return err
	}
	// Encrypt symmetric key with server public key
	encryptedKey, err := crypt.EncryptDataWithServerPubKey(key, serverPublicKey)
	if err != nil {
		return err
	}
	// Create Sirena request
	request := &Request{
		Message: encryptedKey,
	}
	// Set request header
	request.Header = NewHeader(NewHeaderParams{
		Message:    encryptedKey,
		UseEncrypt: true,
	})
	// Set request subheader
	request.SubHeader = MakeSubHeader(encryptedKey)
	clientPrivateKey, err := config.GetKeyFile(config.ClientPrivateKey)
	if err != nil {
		return err
	}
	encryptedKeySignature, err := crypt.GeneratePrivateKeySignature(encryptedKey, clientPrivateKey, config.ClientPrivateKeyPassword)
	if err != nil {
		return err
	}
	// Set request signature
	request.MessageSignature = encryptedKeySignature
	// Send request to Sirena
	response, err := client.Send(request)
	if err != nil {
		return err
	}
	// Validate response header
	if request.Header.ClientID != response.Header.ClientID {
		return fmt.Errorf("request.Header.ClientID (%d) != response.Header.ClientID (%d)", request.Header.ClientID, response.Header.ClientID)
	}
	if request.Header.CreatedAt != response.Header.CreatedAt {
		return fmt.Errorf("request.Header.CreatedAt (%d) != response.Header.CreatedAt (%d)", request.Header.CreatedAt, response.Header.CreatedAt)
	}
	// Decrypt response
	responseKey, err := crypt.DecryptDataWithClientPrivateKey(response.Message[4:132], clientPrivateKey, config.ClientPrivateKeyPassword)
	if err != nil {
		return err
	}
	// Make sure request symmetric key = response symmatric key
	if string(key) != string(responseKey) {
		return fmt.Errorf("Request symmetric key (%s) != response symmetric key(%s)", key, responseKey)
	}
	logger.Debugf("DES key %s signed", responseKey)

	client.Key = responseKey

	return nil
}

// Send sends request to Sirena and returns response
func (client *Client) Send(request *Request) (*Response, error) {
	logger := logger.Get()

	var data []byte
	data = append(data, request.Header.ToBytes()...)
	if len(request.SubHeader) > 0 {
		data = append(data, request.SubHeader...)
	}
	data = append(data, request.Message...)
	if len(request.MessageSignature) > 0 {
		data = append(data, request.MessageSignature...)
	}
	if _, err := client.Conn.Write(data); err != nil {
		logger.Error(err)
		return nil, err
	}
	connReader := bufio.NewReader(client.Conn)
	responseHeaderBytes := make([]byte, 100)
	if _, err := connReader.Read(responseHeaderBytes); err != nil {
		logger.Error(err)
		return nil, err
	}
	responseHeader := ParseHeader(responseHeaderBytes)
	if responseHeader.MessageLength == 0 {
		logger.Errorf("Sirena response header doesn't include messahe length: %s", spew.Sdump(responseHeader))
	}
	responseMessageBytes := make([]byte, responseHeader.MessageLength)
	if _, err := io.ReadFull(connReader, responseMessageBytes); err != nil {
		logger.Error(err)
		return nil, err
	}
	return &Response{
		Header:  &responseHeader,
		Message: responseMessageBytes,
	}, nil
}

// SendXMLRequest send XML request to Sirena and expects XML response
func (client *Client) SendXMLRequest(xmlRequest []byte) ([]byte, error) {
	if len(client.Key) == 0 {
		return nil, errors.New("Client doesn't have symmetric key defined")
	}

	// Kepp key copy in case it's refreshed
	requestKey := make([]byte, len(client.Key))
	copy(requestKey, client.Key)

	xmlCrypted, err := des.Encrypt([]byte(xmlRequest), requestKey)
	if err != nil {
		return nil, err
	}

	// Create Sirena request
	request := &Request{
		Message: xmlCrypted,
	}
	// Set request header
	request.Header = NewHeader(NewHeaderParams{
		Message:      xmlCrypted,
		UseSymmetric: true,
	})

	// Send request to Sirena
	response, err := client.Send(request)
	if err != nil {
		return nil, err
	}

	// Validate response header
	if request.Header.ClientID != response.Header.ClientID {
		return nil, fmt.Errorf("request.Header.ClientID (%d) != response.Header.ClientID (%d)", request.Header.ClientID, response.Header.ClientID)
	}
	if request.Header.CreatedAt != response.Header.CreatedAt {
		return nil, fmt.Errorf("request.Header.CreatedAt (%d) != response.Header.CreatedAt (%d)", request.Header.CreatedAt, response.Header.CreatedAt)
	}
	// Decrypt Sirena response
	xmlResponse, err := des.Decrypt(response.Message, requestKey)
	if err != nil {
		return nil, err
	}

	return xmlResponse, nil
}
