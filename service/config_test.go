package service

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/microparts/logs-go"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
)

var (
	// SignedKey is a signed symmetric key to sign in TestKeyCreate and use in TestAvailability
	sc     configuration.SirenaConfig
	logger *logs.Logger
)

func tearUp() {
	err := godotenv.Load(os.Getenv("ENV_FILE"))
	if err != nil {
		log.Fatal("ErrorResponse loading .env file")
	}

	clientID, _ := String2Uint16(os.Getenv("CLIENT_ID"))
	requestHandlersNum, _ := String2Int32(os.Getenv("REQUEST_HANDLERS"))

	sc = configuration.SirenaConfig{
		ClientID:                 clientID,
		Host:                     os.Getenv("HOST"),
		Port:                     os.Getenv("PORT"),
		ClientPublicKeyFile:      os.Getenv("CLIENT_PUBLIC_KEY"),
		ClientPrivateKeyFile:     os.Getenv("CLIENT_PRIVATE_KEY"),
		ServerPublicKeyFile:      os.Getenv("SERVER_PUBLIC_KEY"),
		ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
		KeysPath:                 os.Getenv("KEYS_PATH"),
		SirenaRequestHandlers:    requestHandlersNum,
		ZippedMessaging:          false,
	}
	lc := &logs.Config{
		Level:  "info",
		Format: "test",
	}
	logger, err = logs.NewLogger(lc)
	if err != nil {
		log.Fatal("ErrorResponse loading .env file")
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
func String2Int32(s string) (uint32, error) {
	b, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(b), nil
}

func TestMain(m *testing.M) {
	tearUp()
	retCode := m.Run()
	os.Exit(retCode)
}
