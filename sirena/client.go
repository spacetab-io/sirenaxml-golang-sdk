package sirena

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	l "github.com/microparts/logs-go"
	"github.com/tmconsulting/sirena-config"

	"github.com/tmconsulting/sirenaxml-golang-sdk/crypt"
	"github.com/tmconsulting/sirenaxml-golang-sdk/des"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/random"

	"github.com/davecgh/go-spew/spew"
)

// Client is a sirena client
type Client struct {
	net.Conn
	Key    []byte
	config *sirenaConfig.SirenaConfig
}

type ClientConfig struct {
	ClientID                 string
	Host                     string
	Port                     string
	ClientPublicKey          string
	ClientPrivateKey         string
	ClientPrivateKeyPassword string
	ServerPublicKey          string
}

// NewClientOptions holds named options for NewClient function
type NewClientOptions struct {
	// Test makes creating and signing symmetric key skipped
	Test bool
}

// NewClient connects to Sirena (if not yet) and returns sirena client singleton
func NewClient(sc *sirenaConfig.SirenaConfig, lc *l.Config, options ...NewClientOptions) *Client {
	err := logs.Init(lc)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.Dial("tcp", sc.GetSirenaAddr())
	if err != nil {
		log.Fatal(err)
	}
	client := &Client{
		Conn:   conn,
		Key:    nil,
		config: sc,
	}
	if len(options) == 0 || !options[0].Test {
		// Create symmetric key
		if err := client.CreateAndSignKey(); err != nil {
			log.Fatal(err)
		}
		// Update key every 1 hour
		go func() {
			for range time.Tick(time.Hour) {
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
	logs.Log.Debug("CreateAndSignKey")
	// Create key as a random string of 8 characters
	var key = []byte(random.String(8))
	logs.Log.Debugf("Trying to sign DES key %s with Sirena", key)
	// Get server public key
	serverPublicKey, err := client.config.GetKeyFile(client.config.ServerPublicKey)
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
	request.Header, err = NewHeader(client.config, NewHeaderParams{
		Message:    encryptedKey,
		UseEncrypt: true,
	})
	if err != nil {
		return err
	}
	// Set request subheader
	request.SubHeader = MakeSubHeader(encryptedKey)
	clientPrivateKey, err := client.config.GetKeyFile(client.config.ClientPrivateKey)
	if err != nil {
		return err
	}
	encryptedKeySignature, err := crypt.GeneratePrivateKeySignature(encryptedKey, clientPrivateKey, client.config.ClientPrivateKeyPassword)
	if err != nil {
		logs.Log.Error(err)
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
	responseKey, err := crypt.DecryptDataWithClientPrivateKey(response.Message[4:132], clientPrivateKey, client.config.ClientPrivateKeyPassword)
	if err != nil {
		return err
	}
	// Make sure request symmetric key = response symmatric key
	if string(key) != string(responseKey) {
		return fmt.Errorf("Request symmetric key (%s) != response symmetric key(%s)", key, responseKey)
	}
	logs.Log.Debugf("DES key %s signed", responseKey)

	client.Key = responseKey

	return nil
}

// Send sends request to Sirena and returns response
func (client *Client) Send(request *Request) (*Response, error) {

	logError := func(err error) error {
		errPrefix := fmt.Sprintf("[SirenaHost:%s SirenaPort:%s SirenaClientID:%s]", client.config.Host, client.config.Port, client.config.ClientID)
		logs.Log.Error(errPrefix + " " + err.Error())
		return err
	}
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
		return nil, logError(err)
	}
	connReader := bufio.NewReader(client.Conn)
	responseHeaderBytes := make([]byte, 100)
	if _, err := connReader.Read(responseHeaderBytes); err != nil {
		return nil, logError(err)
	}
	responseHeader := ParseHeader(responseHeaderBytes)
	if responseHeader.MessageLength == 0 {
		logs.Log.Errorf("Sirena response header doesn't include messahe length: %s", spew.Sdump(responseHeader))
	}
	responseMessageBytes := make([]byte, responseHeader.MessageLength)
	if _, err := io.ReadFull(connReader, responseMessageBytes); err != nil {
		return nil, logError(err)
	}
	return &Response{
		Header:  &responseHeader,
		Message: responseMessageBytes,
	}, nil
}

// SendXMLRequestMaxAttempts defines max attempts to re-dial Sirena API
const MaxReDialAttempts int = 3

// SendXMLRequest send XML request to Sirena and expects XML response
func (client *Client) SendXMLRequest(xmlRequest []byte) ([]byte, error) {
	if len(client.Key) == 0 {
		return nil, errors.New("Client doesn't have symmetric key defined")
	}

	var (
		response                       *Response
		xmlRequestCrypted, xmlResponse []byte
		redialAttempt                  = 0
		err                            error
	)

	for {
		redialAttempt++
		if redialAttempt >= MaxReDialAttempts {
			logs.Log.Debugf("Sirena did't respond after 3 request attempts.")
			break
		}

		// Kepp key copy in case it's refreshed
		requestKey := make([]byte, len(client.Key))
		copy(requestKey, client.Key)

		xmlRequestCrypted, err = des.Encrypt([]byte(xmlRequest), requestKey)
		if err != nil {
			return nil, err
		}

		// Create Sirena request
		request := &Request{
			Message: xmlRequestCrypted,
		}
		// Set request header
		request.Header, err = NewHeader(client.config, NewHeaderParams{
			Message:      xmlRequestCrypted,
			UseSymmetric: true,
		})
		if err != nil {
			logs.Log.Error(err)
			return nil, err
		}

		response, err = client.Send(request)
		if err != nil {
			logs.Log.Error(err)
			if err = client.ReDial(); err != nil {
				log.Fatal(err)
			}
			continue
		}
		// Validate response header
		if request.Header.ClientID != response.Header.ClientID {
			return nil, fmt.Errorf("request.Header.ClientID (%d) != response.Header.ClientID (%d)", request.Header.ClientID, response.Header.ClientID)
		}
		if request.Header.CreatedAt != response.Header.CreatedAt {
			return nil, fmt.Errorf("request.Header.CreatedAt (%d) != response.Header.CreatedAt (%d)", request.Header.CreatedAt, response.Header.CreatedAt)
		}
		// Decrypt Sirena response
		xmlResponse, err = des.Decrypt(response.Message, requestKey)
		if err != nil {
			logs.Log.Error(err)
			if err = client.ReDial(); err != nil {
				log.Fatal(err)
			}
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}

	return xmlResponse, nil
}

// ReDial re-connects to Sirena
func (client *Client) ReDial() error {
	logs.Log.Debugf("Reconnecting to Sirena")
	conn, err := net.Dial("tcp", client.config.GetSirenaAddr())
	if err != nil {
		return err
	}
	client.Conn = conn
	if err := client.CreateAndSignKey(); err != nil {
		return err
	}
	return nil
}
