package sdk

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"time"

	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/crypt"
	"github.com/tmconsulting/sirenaxml-golang-sdk/des"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/random"
)

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
func (client *SirenaClient) CreateAndSignKey() error {
	logs.Log.Debug("CreateAndSignKey")

	// Create key as a random string of 8 characters
	var key = []byte(random.String(8))

	// Create Sirena request
	request, err := client.NewSignRequest(key)
	if err != nil {
		return errors.Wrap(err, "making request error")
	}

	// Send request to Sirena
	response, err := client.Send(request)
	if err != nil {
		return errors.Wrap(err, "sending request error")
	}

	// Validate response header
	if request.Header.ClientID != response.Header.ClientID {
		return errors.Errorf("request.Header.ClientID (%d) != response.Header.ClientID (%d)", request.Header.ClientID, response.Header.ClientID)
	}
	if request.Header.CreatedAt != response.Header.CreatedAt {
		return errors.Errorf("request.Header.CreatedAt (%d) != response.Header.CreatedAt (%d)", request.Header.CreatedAt, response.Header.CreatedAt)
	}

	// Decrypt response
	client.Key, err = crypt.DecryptDataWithClientPrivateKey(response.Message[4:132], client.Config.ClientPrivateKey, client.Config.ClientPrivateKeyPassword)
	if err != nil {
		return errors.Wrap(err, "decrypting data with client private key error")
	}
	client.KeyID = response.Header.KeyID

	// Make sure request symmetric key = response symmatric key
	if string(key) != string(client.Key) {
		return errors.Errorf("Request symmetric key (%s) != response symmetric key(%s)", key, client.Key)
	}

	client.Config.UseSymmetricKeyCrypt = true

	logs.Log.Debugf("DES key %s signed, keyID %d", client.Key, client.KeyID)
	return nil
}

// Send sends request to Sirena and returns response
func (client *SirenaClient) Send(request *Request) (*Response, error) {

	clientParams := fmt.Sprintf("[SirenaHost: %s, SirenaPort: %s, SirenaClientID: %d]:", client.Config.Host, client.Config.Port, client.Config.ClientID) + " %s"

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
		return nil, errors.Wrapf(err, clientParams, "request write error")
	}

	err := client.Conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "SetReadDeadline error")
	}

	connReader := bufio.NewReader(client.Conn)
	responseHeaderBytes := make([]byte, 100)
	if _, err := connReader.Read(responseHeaderBytes); err != nil {
		return nil, errors.Wrapf(err, clientParams, "response header read error")
	}
	responseHeader, err := ParseHeader(responseHeaderBytes)
	if err != nil {
		return nil, errors.Wrapf(err, clientParams, "response header parse error")
	}

	responseMessageBytes := make([]byte, responseHeader.MessageLength)
	if _, err := io.ReadFull(connReader, responseMessageBytes); err != nil {
		return nil, errors.Wrapf(err, clientParams, "response read error")
	}

	if client.Key != nil && responseHeader.Flags.Has(EncryptSymmetric) {
		responseMessageBytes, err = des.Decrypt(responseMessageBytes, client.Key)
		if err != nil {
			return nil, err
		}
	}

	if responseHeader.Flags.Has(ZippedResponse) {
		b := bytes.NewReader(responseMessageBytes)
		z, err := zlib.NewReader(b)
		if err != nil {
			return nil, errors.Wrap(err, "zlib new reader error")
		}
		responseMessageBytes, err = ioutil.ReadAll(z)
		if err != nil {
			return nil, err
		}
		err = z.Close()
		if err != nil {
			return nil, err
		}
	}

	// @TODO обработка ошибок
	//if strings.Contains(string(responseMessageBytes), "error") {
	//	var errResp ErrorResponse
	//	err := xml.Unmarshal(responseMessageBytes, &errResp)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return nil, errors.Errorf("error [code %d]: %s", errResp.Code, errResp.Message)
	//}

	return &Response{
		Header:  responseHeader,
		Message: responseMessageBytes,
	}, nil
}

// SendXMLRequestMaxAttempts defines max attempts to re-dial Sirena API
const MaxReDialAttempts int = 3

// SendXMLRequest send XML request to Sirena and expects XML response
func (client *SirenaClient) SendXMLRequest(xmlRequest []byte) ([]byte, error) {
	if len(client.Key) == 0 {
		return nil, errors.New("SirenaClient doesn't have symmetric key defined")
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
			logs.Log.Warnf("Sirena did't respond after %d request attempts!", redialAttempt)
			break
		}

		// Kepp key copy in case it's refreshed
		requestKey := make([]byte, len(client.Key))
		copy(requestKey, client.Key)

		xmlRequestCrypted, err = des.Encrypt([]byte(xmlRequest), requestKey)
		if err != nil {
			return nil, errors.Wrap(err, "encrypting xmlRequest error")
		}

		// Create Sirena request
		request := &Request{
			Message: xmlRequestCrypted,
			Header: NewHeader(&NewHeaderParams{
				ClientID:      client.Config.ClientID,
				MessageLength: uint32(len(xmlRequestCrypted)),
				UseSymmetric:  true,
			}),
		}

		response, err = client.Send(request)
		if err != nil {
			logs.Log.Error(err)
			if err = client.ReDial(); err != nil {
				return nil, err
			}
			continue
		}
		// Validate response header
		if request.Header.ClientID != response.Header.ClientID {
			return nil, errors.Errorf("request.Header.ClientID (%d) != response.Header.ClientID (%d)", request.Header.ClientID, response.Header.ClientID)
		}
		if request.Header.CreatedAt != response.Header.CreatedAt {
			return nil, errors.Errorf("request.Header.CreatedAt (%d) != response.Header.CreatedAt (%d)", request.Header.CreatedAt, response.Header.CreatedAt)
		}
		// Decrypt Sirena response
		xmlResponse, err = des.Decrypt(response.Message, requestKey)
		if err != nil {
			logs.Log.Error(err)
			if err = client.ReDial(); err != nil {
				return nil, err
			}
			continue
		}
		break
	}

	return xmlResponse, nil
}

// ReDial re-connects to Sirena
func (client *SirenaClient) ReDial() error {
	logs.Log.Debugf("Reconnecting to Sirena")
	conn, err := net.Dial("tcp", client.Config.GetSirenaAddr())
	if err != nil {
		return err
	}
	client.Conn = conn
	if err := client.CreateAndSignKey(); err != nil {
		return err
	}
	return nil
}
