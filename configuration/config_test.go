package sirenaXML

import (
	"net"
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
			Ip:                       os.Getenv("IP"),
			MaxConnections:           requestHandlersNum,
			ClientPublicKey:          os.Getenv("CLIENT_PUBLIC_KEY"),
			ClientPrivateKey:         os.Getenv("CLIENT_PRIVATE_KEY"),
			ServerPublicKey:          os.Getenv("SERVER_PUBLIC_KEY"),
			ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
			ZippedMessaging:          false,
			MaxConnectTries:          1,
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
			Ip:                       os.Getenv("IP"),
			MaxConnections:           requestHandlersNum,
			ClientPrivateKey:         os.Getenv("CLIENT_PRIVATE_KEY"),
			ServerPublicKey:          os.Getenv("SERVER_PUBLIC_KEY"),
			ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
			ZippedMessaging:          false,
			MaxConnectTries:          1,
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
			Ip:                       os.Getenv("IP"),
			MaxConnections:           requestHandlersNum,
			ClientPublicKey:          os.Getenv("CLIENT_PUBLIC_KEY"),
			ServerPublicKey:          os.Getenv("SERVER_PUBLIC_KEY"),
			ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
			ZippedMessaging:          false,
			MaxConnectTries:          1,
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
			Ip:                       os.Getenv("IP"),
			MaxConnections:           requestHandlersNum,
			ClientPublicKey:          os.Getenv("CLIENT_PUBLIC_KEY"),
			ClientPrivateKey:         os.Getenv("CLIENT_PRIVATE_KEY"),
			ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
			ZippedMessaging:          false,
			MaxConnectTries:          1,
		}
		err := sc.PrepareKeys()
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
}

func TestConfig_GetAddr(t *testing.T) {
	ip := net.ParseIP("193.104.87.251")
	expectMaps := map[string]net.TCPAddr{
		EnvLearning:   {IP: ip, Port: 34323},
		EnvTesting:    {IP: ip, Port: 34322},
		EnvProduction: {IP: ip, Port: 34321},
	}
	for env, ipAddress := range expectMaps {
		t.Run("success "+env, func(t *testing.T) {
			sc := Config{Environment: env, Ip: ip.String()}
			addr, err := sc.GetAddr()
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			assert.Equal(t, &ipAddress, addr)
		})
	}

	t.Run("error no ip", func(t *testing.T) {
		sc := Config{Environment: EnvProduction}
		_, err := sc.GetAddr()
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})

	t.Run("error no env", func(t *testing.T) {
		sc := Config{Ip: ip.String()}
		_, err := sc.GetAddr()
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
}
