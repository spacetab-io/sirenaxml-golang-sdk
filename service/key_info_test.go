package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/socket/client"
	"github.com/tmconsulting/sirenaxml-golang-sdk/strings"
	"os"
	s "strings"
	"testing"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/socket"
)

var (
	configKey client.Config
)

func tearUpKeyInfo() {
	clientID, _ := strings.String2Uint16(os.Getenv("CLIENT_ID"))
	//requestHandlersNum, _ := strings.String2Int32(os.Getenv("MAX_CONNECTIONS"))

	configKey = client.Config{
		ClientID:                 clientID,
		Environment:              os.Getenv("ENV"),
		Ip:                       os.Getenv("IP"),
		MaxConnections:           3,
		ClientPublicKey:          os.Getenv("CLIENT_PUBLIC_KEY"),
		ClientPrivateKey:         os.Getenv("CLIENT_PRIVATE_KEY"),
		ServerPublicKey:          os.Getenv("SERVER_PUBLIC_KEY"),
		ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
		ZippedMessaging:          false,
		MaxConnectTries:          3,
	}
}

func TestService_KeyInfo(t *testing.T) {
	logger := logs.NewNullLog()
	tearUpKeyInfo()

	configKey.ServerPublicKey = s.ReplaceAll(conf.ServerPublicKey, "\\n", "\n")
	configKey.ClientPublicKey = s.ReplaceAll(conf.ClientPublicKey, "\\n", "\n")
	configKey.ClientPrivateKey = s.ReplaceAll(conf.ClientPrivateKey, "\\n", "\n")

	sdkClient, err := socket.NewClient(
		logger,
		configKey.ClientPrivateKey,
		configKey.ClientPrivateKeyPassword,
		configKey.ClientPublicKey,
		configKey.Ip,
		configKey.Environment,
		configKey.ServerPublicKey,
		configKey.Address,
		configKey.Buffer,
		configKey.ZippedMessaging,
		configKey.MaxConnections,
		configKey.ClientID,
		configKey.MaxConnectTries,
	)

	if !assert.NoError(t, err) {
		t.FailNow()
	}

	service := NewSKD(sdkClient)
	checkKeyData(t, sdkClient)
	t.Run("success", func(t *testing.T) {
		_, err = service.KeyInfo()
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})
}

func checkKeyData(t *testing.T, c Storage) {
	kd, err := c.GetKeyData()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !assert.NotEmpty(t, kd.Key) {
		t.FailNow()
	}
	if !assert.NotZero(t, kd.ID) {
		t.FailNow()
	}
}
