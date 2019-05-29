package service

import (
	"os"
	"strconv"
	"testing"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
)

var (
	sc sirena.Config
)

func tearUp() {
	clientID, _ := String2Uint16(os.Getenv("CLIENT_ID"))
	requestHandlersNum, _ := String2Int32(os.Getenv("REQUEST_HANDLERS"))

	sc = sirena.Config{
		ClientID:                 clientID,
		Ip:                       os.Getenv("HOST"),
		Port:                     os.Getenv("PORT"),
		ClientPublicKey:          []byte(os.Getenv("CLIENT_PUBLIC_KEY")),
		ClientPrivateKey:         []byte(os.Getenv("CLIENT_PRIVATE_KEY")),
		ServerPublicKey:          []byte(os.Getenv("SERVER_PUBLIC_KEY")),
		ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
		KeysPath:                 os.Getenv("KEYS_PATH"),
		RequestHandlers:          requestHandlersNum,
		ZippedMessaging:          false,
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
