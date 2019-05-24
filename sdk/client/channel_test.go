package client

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/microparts/logs-go"
	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
)

var (
	// SignedKey is a signed symmetric key to sign in TestKeyCreate and use in TestAvailability
	sc *configuration.SirenaConfig
	lc *logs.Config
)

func tearUp() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("loading .env file error: %v", err)
	}

	clientID, _ := String2Uint16(os.Getenv("CLIENT_ID"))
	handlersNum, _ := String2Uint32(os.Getenv("REQUEST_HANDLERS"))
	sc = &configuration.SirenaConfig{
		ClientID:                 clientID,
		Host:                     os.Getenv("HOST"),
		Port:                     os.Getenv("PORT"),
		SirenaRequestHandlers:    handlersNum,
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
		Level:  "info",
		Format: "test",
	}
}

// String2Uint16 converts string to uint16
func String2Uint16(s string) (uint16, error) {
	b, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(b), nil
}

// String2Uint16 converts string to uint16
func String2Uint32(s string) (uint32, error) {
	b, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(b), nil
}

//func TestMain(m *testing.M) {
//	tearUp()
//	retCode := m.Run()
//	os.Exit(retCode)
//}

func TestNewChannel(t *testing.T) {
	tearUp()

	c, err := NewChannel(sc)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	if !assert.NotEmpty(t, c.KeyID) {
		t.FailNow()
	}
}
