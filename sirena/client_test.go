package sirena

import (
	"encoding/xml"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/microparts/logs-go"
	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/utils"
)

var (
	// SignedKey is a signed symmetric key to sign in TestKeyCreate and use in TestAvailability
	sc *configuration.SirenaConfig
	lc *logs.Config
)

var keyInfoXML = []byte(`<?xml version="1.0" encoding="UTF-8"?><sirena><query><key_info/></query></sirena>`)

func tearUp() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("ErrorResponse loading .env file")
	}

	clientID, _ := utils.String2Uint16(os.Getenv("CLIENT_ID"))

	sc = &configuration.SirenaConfig{
		ClientID:                 clientID,
		Host:                     os.Getenv("HOST"),
		Port:                     os.Getenv("PORT"),
		ClientPublicKeyFile:      os.Getenv("CLIENT_PUBLIC_KEY"),
		ClientPrivateKeyFile:     os.Getenv("CLIENT_PRIVATE_KEY"),
		ServerPublicKeyFile:      os.Getenv("SERVER_PUBLIC_KEY"),
		ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
		KeysPath:                 os.Getenv("KEYS_PATH"),
		UseSymmetricKeyCrypt:     false,
		ZipRequests:              false,
		ZipResponses:             false,
	}
	lc = &logs.Config{
		Level:  "debug",
		Format: "test",
	}
}

func TestMain(m *testing.M) {
	tearUp()
	retCode := m.Run()
	os.Exit(retCode)
}

// @TODO переделать
func TestKeyInfo(t *testing.T) {
	client, err := NewClient(sc, lc)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	request := &Request{
		Message: keyInfoXML,
		Header: NewHeader(&NewHeaderParams{
			ClientID:      client.Config.ClientID,
			MessageLength: uint32(len(keyInfoXML)),
			//CanBeZipped:   true,
		}),
	}
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	response, err := client.Send(request)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	// Validate response header
	assert.Equal(t, request.Header.Flags, response.Header.Flags)
	assert.Equal(t, request.Header.ClientID, response.Header.ClientID)
	assert.Equal(t, request.Header.CreatedAt, response.Header.CreatedAt)

	// Decode XML and make sure Sirena Public Key is returned
	var keyInfoResponse = struct {
		Answer struct {
			KeyInfo struct {
				KeyManager struct {
					ServerPubliKey string `xml:"server_public_key"`
				} `xml:"key_manager"`
			} `xml:"key_info"`
		} `xml:"answer"`
	}{}
	if err := xml.Unmarshal(response.Message, &keyInfoResponse); err != nil {
		t.Fatalf("unmarshall message %s error: %v", string(response.Message), err)
	}
	assert.NotEmpty(t, keyInfoResponse.Answer.KeyInfo.KeyManager.ServerPubliKey)
}

// @TODO переделать
//func TestKeyCreate(t *testing.T) {
//	t.Run("test signing DES key", func(t *testing.T) {
//		client, err := NewClient(sc, lc)
//		if !assert.NoError(t, err) {
//			t.FailNow()
//		}
//
//		// Create Sirena request
//		request := &Request{
//			Message: encryptedKey,
//			Header: NewHeader(&NewHeaderParams{
//				ClientID:      client.Config.ClientID,
//				MessageLength: uint32(len(encryptedKey)),
//				UseEncrypt:    true,
//				//CanBeZipped:   true,
//				//UseSymmetric:  true,
//			}),
//		}
//		// Set request header
//		if !assert.NoError(t, err) {
//			t.FailNow()
//		}
//		// Set request subheader
//		request.SubHeader = MakeSubHeader(encryptedKey)
//
//		encryptedKeySignature, err := crypt.GeneratePrivateKeySignature(encryptedKey, client.Config.ClientPrivateKey, client.Config.ClientPrivateKeyPassword)
//		if !assert.NoError(t, err) {
//			t.FailNow()
//		}
//		// Set request signature
//		request.MessageSignature = encryptedKeySignature
//		// Send request to Sirena
//		response, err := client.Send(request)
//		if !assert.NoError(t, err) {
//			t.FailNow()
//		}
//		// Validate response header
//		assert.Equal(t, request.Header.ClientID, response.Header.ClientID)
//		assert.Equal(t, request.Header.CreatedAt, response.Header.CreatedAt)
//
//		// Decrypt response
//		responseKey, err := crypt.DecryptDataWithClientPrivateKey(response.Message[4:132], client.Config.ClientPrivateKey, client.Config.ClientPrivateKeyPassword)
//		if !assert.NoError(t, err) {
//			t.FailNow()
//		}
//		// Make sure request symmetric key = response symmatric key
//		assert.Equal(t, string(key), string(responseKey))
//
//		SignedKey = responseKey
//	})
//}

// AvailabilityXML is a test availability XML
func TestClient_Send(t *testing.T) {
	t.Run("test zipped request", func(t *testing.T) {
		customSirenConfig := sc
		customSirenConfig.ZipRequests = true
		testAvailability(t, customSirenConfig)
	})
	t.Run("test zipped request and response", func(t *testing.T) {
		customSirenConfig := sc
		customSirenConfig.ZipRequests = true
		customSirenConfig.ZipResponses = true
		testAvailability(t, customSirenConfig)
	})
	t.Run("test symmetric key encription", func(t *testing.T) {
		customSirenConfig := sc
		customSirenConfig.UseSymmetricKeyCrypt = true
		testAvailability(t, customSirenConfig)
	})
	t.Run("test all flags", func(t *testing.T) {
		customSirenConfig := sc
		customSirenConfig.ZipRequests = true
		customSirenConfig.ZipResponses = true
		customSirenConfig.UseSymmetricKeyCrypt = true
		testAvailability(t, customSirenConfig)
	})
}

func testAvailability(t *testing.T, sc *configuration.SirenaConfig) {
	client, err := NewClient(sc, lc)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !assert.NotEmpty(t, client.Key) {
		t.FailNow()
	}
	availabiliteReq := &AvailabilityRequest{
		Query: AvailabilityRequestQuery{
			Availability: Availability{
				Departure: "MOW",
				Arrival:   "LED",
				AnswerParams: AvailabilityAnswerParams{
					ShowFlighttime: true,
				},
			},
		},
	}
	availabiliteReqXML, err := xml.Marshal(availabiliteReq)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	// Create Sirena request
	request, err := client.NewRequest(availabiliteReqXML)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	// Send request to Sirena
	response, err := client.Send(request)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	// Validate response header
	assert.Equal(t, request.Header.ClientID, response.Header.ClientID)
	assert.Equal(t, request.Header.CreatedAt, response.Header.CreatedAt)
	// Decode XML and make sure availability returned
	var availabilityResponse AvailabilityResponse
	if err := xml.Unmarshal(response.Message, &availabilityResponse); err != nil {
		t.Fatalf("unmarshall message %s error: %v", string(response.Message), err)
	}
	//log.Print(string(response.Message))
	//log.Printf("availabilityResponse: %+v", availabilityResponse)
	// Check Sirena availability response
	assert.NotEmpty(t, availabilityResponse.Answer.Availability.Flights)
}
