package sirenaXML

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/strings"
)

func TestConfig_PrepareKeys(t *testing.T) {
	clientID, _ := strings.String2Uint16(os.Getenv("CLIENT_ID"))
	requestHandlersNum, _ := strings.String2Int32(os.Getenv("MAX_CONNECTIONS"))

	t.Run("good keys", func(t *testing.T) {
		sc := Config{
			ClientID:                 clientID,
			Environment:              os.Getenv("ENV"),
			MaxConnections:           requestHandlersNum,
			ClientPublicKey:          os.Getenv("CLIENT_PUBLIC_KEY"),
			ClientPrivateKey:         os.Getenv("CLIENT_PRIVATE_KEY"),
			ServerPublicKey:          os.Getenv("SERVER_PUBLIC_KEY"),
			ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
			ZippedMessaging:          false,
		}
		err := sc.PrepareKeys()
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})

	t.Run("empty ClientPublicKey", func(t *testing.T) {
		sc := Config{
			ClientID:                 clientID,
			Environment:              os.Getenv("ENV"),
			MaxConnections:           requestHandlersNum,
			ClientPrivateKey:         os.Getenv("CLIENT_PRIVATE_KEY"),
			ServerPublicKey:          os.Getenv("SERVER_PUBLIC_KEY"),
			ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
			ZippedMessaging:          false,
		}
		err := sc.PrepareKeys()
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})

	t.Run("empty ClientPrivateKey", func(t *testing.T) {
		sc := Config{
			ClientID:                 clientID,
			Environment:              os.Getenv("ENV"),
			MaxConnections:           requestHandlersNum,
			ClientPublicKey:          os.Getenv("CLIENT_PUBLIC_KEY"),
			ServerPublicKey:          os.Getenv("SERVER_PUBLIC_KEY"),
			ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
			ZippedMessaging:          false,
		}
		err := sc.PrepareKeys()
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})

	t.Run("empty ServerPublicKey", func(t *testing.T) {
		sc := Config{
			ClientID:                 clientID,
			Environment:              os.Getenv("ENV"),
			MaxConnections:           requestHandlersNum,
			ClientPublicKey:          os.Getenv("CLIENT_PUBLIC_KEY"),
			ClientPrivateKey:         os.Getenv("CLIENT_PRIVATE_KEY"),
			ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
			ZippedMessaging:          false,
		}
		err := sc.PrepareKeys()
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
}

func TestConfig_GetAddr(t *testing.T) {
	expectMaps := map[string][]string{
		EnvLearning:   {"193.104.87.251:34323", "194.84.25.50:34323"},
		EnvTesting:    {"193.104.87.251:34322", "194.84.25.50:34322"},
		EnvProduction: {"193.104.87.251:34321", "194.84.25.50:34321"},
	}
	for env, slice := range expectMaps {
		t.Run("success "+env, func(t *testing.T) {
			sc := Config{Environment: env}
			addr, err := sc.GetAddr()
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			assert.Subset(t, slice, []string{addr})
		})
	}

	t.Run("error", func(t *testing.T) {
		sc := Config{}
		_, err := sc.GetAddr()
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
}
