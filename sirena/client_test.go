package sirena

import (
	"encoding/xml"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/microparts/logs-go"
	"github.com/stretchr/testify/assert"
	"github.com/tmconsulting/sirena-config"

	"github.com/tmconsulting/sirenaxml-golang-sdk/crypt"
	"github.com/tmconsulting/sirenaxml-golang-sdk/des"
	"github.com/tmconsulting/sirenaxml-golang-sdk/random"
)

// SignedKey is a signed symmetric key to sigin in TestKeyCreate and use in TestAvailability
var (
	SignedKey []byte
	sc        *sirenaConfig.SirenaConfig
	lc        *logs.Config
)

const keyInfoXML = `<?xml version="1.0" encoding="UTF-8"?>
<sirena>
  <query>
    <key_info/>
  </query>
</sirena>`

func tearUp() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sc = &sirenaConfig.SirenaConfig{
		ClientID:                 os.Getenv("CLIENT_ID"),
		Host:                     os.Getenv("HOST"),
		Port:                     os.Getenv("PORT"),
		ClientPublicKey:          os.Getenv("CLIENT_PUBLIC_KEY"),
		ClientPrivateKey:         os.Getenv("CLIENT_PRIVATE_KEY"),
		ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
		ServerPublicKey:          os.Getenv("SERVER_PUBLIC_KEY"),
		KeysPath:                 os.Getenv("KEYS_PATH"),
	}
	lc = &logs.Config{
		Level:  "info",
		Format: "json",
	}
}

func TestMain(m *testing.M) {
	tearUp()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestKeyInfo(t *testing.T) {
	client, err := NewClient(sc, lc, NewClientOptions{Test: true})
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	request := &Request{
		Message: []byte(keyInfoXML),
	}
	request.Header, err = NewHeader(client.config, NewHeaderParams{
		Message: request.Message,
	})
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	response, err := client.Send(request)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	// Validate response header
	if request.Header.ClientID != response.Header.ClientID {
		t.Fatalf("request.Header.ClientID (%d) != response.Header.ClientID (%d)", request.Header.ClientID, response.Header.ClientID)
	}
	if request.Header.CreatedAt != response.Header.CreatedAt {
		t.Fatalf("request.Header.CreatedAt (%d) != response.Header.CreatedAt (%d)", request.Header.CreatedAt, response.Header.CreatedAt)
	}
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
		t.Fatal(err)
	}
	if keyInfoResponse.Answer.KeyInfo.KeyManager.ServerPubliKey == "" {
		t.Fatalf("No Sirena Public key found in response: %s", response.Message)
	}
	t.Logf("Got Sirena public key: \n%s", keyInfoResponse.Answer.KeyInfo.KeyManager.ServerPubliKey)
}

func TestKeyCreate(t *testing.T) {
	client, err := NewClient(sc, lc, NewClientOptions{Test: true})
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	// Create symmetric key
	var key = []byte(random.String(8))
	t.Logf("Trying to sign DES key %s with Sirena", key)
	// Get server public key
	serverPublicKey, err := sc.GetKeyFile(sc.ServerPublicKey)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	// Encrypt symmetric key with server public key
	encryptedKey, err := crypt.EncryptDataWithServerPubKey(key, serverPublicKey)
	if !assert.NoError(t, err) {
		t.FailNow()
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
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	// Set request subheader
	request.SubHeader = MakeSubHeader(encryptedKey)
	clientPrivateKey, err := sc.GetKeyFile(sc.ClientPrivateKey)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	encryptedKeySignature, err := crypt.GeneratePrivateKeySignature(encryptedKey, clientPrivateKey, sc.ClientPrivateKeyPassword)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	// Set request signature
	request.MessageSignature = encryptedKeySignature
	// Send request to Sirena
	response, err := client.Send(request)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	// Validate response header
	if request.Header.ClientID != response.Header.ClientID {
		t.Errorf("request.Header.ClientID (%d) != response.Header.ClientID (%d)", request.Header.ClientID, response.Header.ClientID)
	}
	if request.Header.CreatedAt != response.Header.CreatedAt {
		t.Errorf("request.Header.CreatedAt (%d) != response.Header.CreatedAt (%d)", request.Header.CreatedAt, response.Header.CreatedAt)
	}
	// Decrypt response
	responseKey, err := crypt.DecryptDataWithClientPrivateKey(response.Message[4:132], clientPrivateKey, sc.ClientPrivateKeyPassword)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	// Make sure request symmetric key = response symmatric key
	if string(key) != string(responseKey) {
		t.Errorf("Request symmetric key (%s) != response symmetric key(%s)", key, responseKey)
	}
	t.Logf("DES key %s signed", responseKey)
	SignedKey = responseKey
}

// AvailabilityXML is a test availability XML
const AvailabilityXML = `<?xml version="1.0" encoding="UTF-8"?>
<sirena>
<query>
  <availability>
    <departure>МОВ</departure>
    <arrival>ХБР</arrival>
    <answer_params>
      <show_flighttime>true</show_flighttime>
    </answer_params>
  </availability>
</query>
</sirena>`

func TestAvailability(t *testing.T) {
	client, err := NewClient(sc, lc, NewClientOptions{Test: true})
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	if len(SignedKey) == 0 {
		t.Fatal("No signed key found")
	}

	xmlCrypted, err := des.Encrypt([]byte(AvailabilityXML), SignedKey)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	// Create Sirena request
	request := &Request{
		Message: xmlCrypted,
	}
	// Set request header
	request.Header, err = NewHeader(client.config, NewHeaderParams{
		Message:      xmlCrypted,
		UseSymmetric: true,
	})
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	// Send request to Sirena
	response, err := client.Send(request)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	// Validate response header
	if request.Header.ClientID != response.Header.ClientID {
		t.Errorf("request.Header.ClientID (%d) != response.Header.ClientID (%d)", request.Header.ClientID, response.Header.ClientID)
	}
	if request.Header.CreatedAt != response.Header.CreatedAt {
		t.Errorf("request.Header.CreatedAt (%d) != response.Header.CreatedAt (%d)", request.Header.CreatedAt, response.Header.CreatedAt)
	}
	// Decrypt Sirena response
	xmlResponse, err := des.Decrypt(response.Message, SignedKey)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	// Decode XML and make sure availability returned
	var availabilityResponse = struct {
		Answer struct {
			Availability struct {
				Departure string `xml:"departure,attr"`
				Arrival   string `xml:"arrival,attr"`
			} `xml:"availability"`
		} `xml:"answer"`
	}{}
	if err := xml.Unmarshal(xmlResponse, &availabilityResponse); err != nil {
		t.Fatal(err)
	}
	if availabilityResponse.Answer.Availability.Arrival == "" || availabilityResponse.Answer.Availability.Departure == "" {
		t.Fatalf("Invalid Sirena availability response: %s", xmlResponse)
	}
	t.Logf("Availability response is correct:\n%s", xmlResponse)
}
