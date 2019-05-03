package client

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"net"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
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
	try              int
}

type RequestError struct {
	MessageID uint32
	Header    *Header
	Code      string
	Message   string
	Error     error
}

// Response is a Sirena response
type Response struct {
	Header    *Header
	MessageID uint32
	Message   []byte
}

// SirenaClient is a sirena client
type SirenaClient struct {
	net.Conn
	Key            []byte
	KeyID          uint32
	Config         *configuration.SirenaConfig
	respChan       map[uint32]chan *Response
	errChan        map[uint32]chan *RequestError
	requests       map[uint32]*Request
	msgPool        *msgIDsPool
	socketWriteMux sync.Mutex
	respMux        sync.RWMutex
	reqMux         sync.Mutex
}

const ReadingSocketError = 999
const RequestTimeout = 5

var (
	errorFormat         string
	RequestTimeoutError = errors.New("request timeout")
)

// NewSirenaClient connects to Sirena (if not yet) and returns sirena client singleton
func NewSirenaClient(sc *configuration.SirenaConfig) (*SirenaClient, error) {
	err := sc.GetCerts()
	if err != nil {
		return nil, err
	}
	conn, err := net.Dial("tcp", sc.GetSirenaAddr())
	if err != nil {
		return nil, errors.Wrap(err, "dial sirena addr error")
	}
	err = conn.SetDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		return nil, err

	}
	client := &SirenaClient{
		Conn:     conn,
		msgPool:  makeMsgPool(),
		Key:      nil,
		Config:   sc,
		requests: make(map[uint32]*Request),
		respChan: make(map[uint32]chan *Response, 1),
	}

	if client.Key == nil {
		// Create symmetric key
		if err := client.CreateAndSignKey(); err != nil {
			return nil, errors.Wrap(err, "creating and signing key error")
		}
		// Update key every 1 hour
		// @TODO что-то это нихера не очевидная ебулдень. Оно точно будет работать и будет работать корректно?
		go func() {
			for range time.Tick(time.Hour) {
				if err := client.CreateAndSignKey(); err != nil {
					logs.Log.Fatal("key updating error")
				}
			}
		}()
	}
	go client.listenSocketContinuously()
	errorFormat = fmt.Sprintf("[SirenaHost: %s, SirenaPort: %s, SirenaClientID: %d]:", client.Config.Host, client.Config.Port, client.Config.ClientID) + " %s"
	return client, nil
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
	client.addRequestToQueue(request)
	err := client.sendToSocket(request)
	if err != nil {
		return nil, err
	}

	if client.Key == nil {
		go client.getMessageFromSocket(-1)
	}

	var response *Response
	select {
	case <-time.After(RequestTimeout * time.Second):
		client.removeRequestFromQueue(request.Header.MessageID)
		logs.Log.Debugf("[%d] response not received, due timeout", request.Header.MessageID)
		return nil, RequestTimeoutError
	case response = <-client.respChan[request.Header.MessageID]:
		client.msgPool.returnMsgIDToPool(request.Header.MessageID)
		client.removeRequestFromQueue(request.Header.MessageID)
		logs.Log.Debugf("[%d] response received, msgId returned", request.Header.MessageID)
	}

	responseMessageBytes := response.Message
	if client.Key != nil && response.Header.Flags.Has(EncryptSymmetric) {
		responseMessageBytes, err = des.Decrypt(responseMessageBytes, client.Key)
		if err != nil {
			return nil, err
		}
	}

	if response.Header.Flags.Has(ZippedResponse) {
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
		Header:  response.Header,
		Message: responseMessageBytes,
	}, nil
}

func (client *SirenaClient) resendRequest(msgID uint32) {
	if client.requests[msgID].try == MaxTries {
		return
	}
	client.requests[msgID].try++
	_ = client.sendToSocket(client.requests[msgID])
	return
}

func (client *SirenaClient) addRequestToQueue(request *Request) {
	client.reqMux.Lock()
	client.requests[request.Header.MessageID] = request
	client.reqMux.Unlock()
}

func (client *SirenaClient) removeRequestFromQueue(msgID uint32) {
	client.reqMux.Lock()
	delete(client.requests, msgID)
	client.reqMux.Unlock()
}

// SendXMLRequestMaxAttempts defines max attempts to re-dial Sirena API
//const MaxReDialAttempts int = 3

// SendXMLRequest send XML request to Sirena and expects XML response
// Deprecated
//func (client *SirenaClient) SendXMLRequest(xmlRequest []byte) ([]byte, error) {
//	if len(client.Key) == 0 {
//		return nil, errors.New("SirenaClient doesn't have symmetric key defined")
//	}
//
//	var (
//		response                       *Response
//		xmlRequestCrypted, xmlResponse []byte
//		redialAttempt                  = 0
//		err                            error
//	)
//
//	for {
//		redialAttempt++
//		if redialAttempt >= MaxReDialAttempts {
//			logs.Log.Warnf("Sirena did't respond after %d request attempts!", redialAttempt)
//			break
//		}
//
//		// Kepp key copy in case it's refreshed
//		requestKey := make([]byte, len(client.Key))
//		copy(requestKey, client.Key)
//
//		xmlRequestCrypted, err = des.Encrypt([]byte(xmlRequest), requestKey)
//		if err != nil {
//			return nil, errors.Wrap(err, "encrypting xmlRequest error")
//		}
//
//		// Create Sirena request
//		request := &Request{
//			Message: xmlRequestCrypted,
//			Header: NewHeader(&NewHeaderParams{
//				ClientID:      client.Config.ClientID,
//				MessageLength: uint32(len(xmlRequestCrypted)),
//				UseSymmetric:  true,
//			}),
//		}
//
//		response, err = client.Send(request)
//		if err != nil {
//			logs.Log.Error(err)
//			if err = client.ReDial(); err != nil {
//				return nil, err
//			}
//			continue
//		}
//		// Validate response header
//		if request.Header.ClientID != response.Header.ClientID {
//			return nil, errors.Errorf("request.Header.ClientID (%d) != response.Header.ClientID (%d)", request.Header.ClientID, response.Header.ClientID)
//		}
//		if request.Header.CreatedAt != response.Header.CreatedAt {
//			return nil, errors.Errorf("request.Header.CreatedAt (%d) != response.Header.CreatedAt (%d)", request.Header.CreatedAt, response.Header.CreatedAt)
//		}
//		// Decrypt Sirena response
//		xmlResponse, err = des.Decrypt(response.Message, requestKey)
//		if err != nil {
//			logs.Log.Error(err)
//			if err = client.ReDial(); err != nil {
//				return nil, err
//			}
//			continue
//		}
//		break
//	}
//
//	return xmlResponse, nil
//}

// ReDial re-connects to Sirena
//func (client *SirenaClient) ReDial() error {
//	logs.Log.Debugf("Reconnecting to Sirena")
//	conn, err := net.Dial("tcp", client.Config.GetSirenaAddr())
//	if err != nil {
//		return err
//	}
//	client.Conn = conn
//	if err := client.CreateAndSignKey(); err != nil {
//		return err
//	}
//	return nil
//}
