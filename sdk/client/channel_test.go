package client

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
)

var (
	sc *sirenaXML.Config
)

func tearUp() {
	clientID, _ := String2Uint16(os.Getenv("CLIENT_ID"))
	handlersNum, _ := String2Uint32(os.Getenv("MAX_CONNECTIONS"))
	sc = &sirenaXML.Config{
		ClientID:                 clientID,
		Environment:              os.Getenv("ENV"),
		Ip:                       os.Getenv("IP"),
		MaxConnections:           handlersNum,
		ClientPublicKey:          os.Getenv("CLIENT_PUBLIC_KEY"),
		ClientPrivateKey:         os.Getenv("CLIENT_PRIVATE_KEY"),
		ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
		ServerPublicKey:          os.Getenv("SERVER_PUBLIC_KEY"),
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
func String2Uint32(s string) (uint32, error) {
	b, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(b), nil
}

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
