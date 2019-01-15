package sirena

import (
	"encoding/xml"
	"testing"

	"github.com/tmconsulting/sirenaxml-golang-sdk/random"

	"gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/sirena-agent-go/config"

	"github.com/tmconsulting/sirenaxml-golang-sdk/des"

	"github.com/tmconsulting/sirenaxml-golang-sdk/crypt"
)

// SignedKey is a signed symmetric key to sigin in TestKeyCreate and use in TestAvailability
var SignedKey []byte

const keyInfoXML = `<?xml version="1.0" encoding="UTF-8"?>
<sirena>
  <query>
    <key_info/>
  </query>
</sirena>`

func TestKeyInfo(t *testing.T) {
	client := NewClient(NewClientOptions{Test: true})
	request := &Request{
		Message: []byte(keyInfoXML),
	}
	request.Header = NewHeader(NewHeaderParams{
		Message: request.Message,
	})
	response, err := client.Send(request)
	if err != nil {
		t.Fatal(err)
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
	client := NewClient(NewClientOptions{Test: true})
	config := config.Get()
	// Create symmetric key
	var key = []byte(random.String(8))
	t.Logf("Trying to sign DES key %s with Sirena", key)
	// Get server public key
	serverPublicKey, err := config.GetKeyFile(config.ServerPublicKey)
	if err != nil {
		t.Fatal(err)
	}
	// Encrypt symmetric key with server public key
	encryptedKey, err := crypt.EncryptDataWithServerPubKey(key, serverPublicKey)
	if err != nil {
		t.Fatal(err)
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
		t.Fatal(err)
	}
	encryptedKeySignature, err := crypt.GeneratePrivateKeySignature(encryptedKey, clientPrivateKey, config.ClientPrivateKeyPassword)
	if err != nil {
		t.Fatal(err)
	}
	// Set request signature
	request.MessageSignature = encryptedKeySignature
	// Send request to Sirena
	response, err := client.Send(request)
	if err != nil {
		t.Fatal(err)
	}
	// Validate response header
	if request.Header.ClientID != response.Header.ClientID {
		t.Errorf("request.Header.ClientID (%d) != response.Header.ClientID (%d)", request.Header.ClientID, response.Header.ClientID)
	}
	if request.Header.CreatedAt != response.Header.CreatedAt {
		t.Errorf("request.Header.CreatedAt (%d) != response.Header.CreatedAt (%d)", request.Header.CreatedAt, response.Header.CreatedAt)
	}
	// Decrypt response
	responseKey, err := crypt.DecryptDataWithClientPrivateKey(response.Message[4:132], clientPrivateKey, config.ClientPrivateKeyPassword)
	if err != nil {
		t.Fatal(err)
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
	client := NewClient(NewClientOptions{Test: true})

	if len(SignedKey) == 0 {
		t.Fatal("No signed key found")
	}

	xmlCrypted, err := des.Encrypt([]byte(AvailabilityXML), SignedKey)
	if err != nil {
		t.Fatal(err)
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
		t.Fatal(err)
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
	if err != nil {
		t.Fatal(err)
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
