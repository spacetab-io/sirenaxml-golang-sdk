package client

import (
	"os"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
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
	}
}

func TestNewChannel(t *testing.T) {
	tearUp()
	l := logs.NewNullLog()
	t.Run("success", func(t *testing.T) {
		c, err := NewChannel(&sc, l)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		if !assert.NotEmpty(t, c.socket.KeyData.ID) {
			t.FailNow()
		}
		assert.NotPanics(t, func() {
			c.stopListener()
		})
	})
	t.Run("reconnect", func(t *testing.T) {
		c, err := NewChannel(&sc, l)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		if !assert.NotEmpty(t, c.socket.KeyData.ID) {
			t.FailNow()
		}
		time.Sleep(1 * time.Second)
		assert.NotPanics(t, func() {
			c.stopListener()
		})
	})
	t.Run("error throttling", func(t *testing.T) {
		c, err := NewChannel(&sc, l)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.Panics(t, func() {
			c.reconnect(errors.New("throttling"))
		})
		assert.NotPanics(t, func() {
			c.stopListener()
		})
	})
}
