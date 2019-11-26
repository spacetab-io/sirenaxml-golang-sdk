package service

import (
	"os"
	"testing"

	"github.com/tmconsulting/sirenaxml-golang-sdk/strings"
)

var (
	sc sirenaXML.Config
)

func tearUp() {
	clientID, _ := strings.String2Uint16(os.Getenv("CLIENT_ID"))
	requestHandlersNum, _ := strings.String2Int32(os.Getenv("MAX_CONNECTIONS"))

	sc = sirenaXML.Config{
		ClientID:                 clientID,
		Environment:              os.Getenv("ENV"),
		Ip:                       os.Getenv("IP"),
		MaxConnections:           requestHandlersNum,
		ClientPublicKey:          os.Getenv("CLIENT_PUBLIC_KEY"),
		ClientPrivateKey:         os.Getenv("CLIENT_PRIVATE_KEY"),
		ServerPublicKey:          os.Getenv("SERVER_PUBLIC_KEY"),
		ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
		ZippedMessaging:          false,
		MaxConnectTries:          3,
	}
}

func TestMain(m *testing.M) {
	tearUp()
	retCode := m.Run()
	os.Exit(retCode)
}
