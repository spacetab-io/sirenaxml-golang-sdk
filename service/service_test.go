package service

import (
	"encoding/xml"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk/client"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

// AvailabilityXML is a test availability XML
func TestService(t *testing.T) {
	t.Run("test no zipped request", func(t *testing.T) {
		customSirenConfig := conf
		testRequest(t, customSirenConfig)
	})
	t.Run("test zipped request/response", func(t *testing.T) {
		customSirenConfig := conf
		customSirenConfig.ZippedMessaging = true
		testRequest(t, customSirenConfig)
	})
	t.Run("test error params", func(t *testing.T) {
		logger := logs.NewNullLog()
		customSirenConfig := conf
		customSirenConfig.ClientID = 1111

		_, err := sdk.NewClient(
			logger,
			customSirenConfig.ClientPrivateKey,
			customSirenConfig.ClientPrivateKeyPassword,
			customSirenConfig.ClientPublicKey,
			customSirenConfig.Ip, conf.Environment,
			customSirenConfig.ServerPublicKey,
			customSirenConfig.Address,
			customSirenConfig.Buffer,
			customSirenConfig.ZippedMessaging,
			customSirenConfig.MaxConnections,
			customSirenConfig.ClientID,
		)

		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
}

func TestService_RawRequest(t *testing.T) {
	logger := logs.NewNullLog()
	sdkClient, err := sdk.NewClient(
		logger,
		conf.ClientPrivateKey,
		conf.ClientPrivateKeyPassword,
		conf.ClientPublicKey,
		conf.Ip,
		conf.Environment,
		conf.ServerPublicKey,
		conf.Address,
		conf.Buffer,
		conf.ZippedMessaging,
		conf.MaxConnections,
		conf.ClientID,
	)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	service := NewSKD(sdkClient)
	checkKeyData(t, sdkClient)
	t.Run("success", func(t *testing.T) {
		xmlReq := []byte(`<?xml version="1.0" encoding="UTF-8"?><sirena><query><key_info/></query></sirena>`)
		response, err := service.RawRequest(xmlReq)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		// Decode XML and make sure Sirena Public Key is returned
		var keyInfoResponse structs.KeyInfoResponse
		err = xml.Unmarshal(response, &keyInfoResponse)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		if !assert.NotEmpty(t, keyInfoResponse.Answer.KeyInfo.KeyManager.ServerPublicKey) {
			t.FailNow()
		}
	})
}

func testRequest(t *testing.T, sc client.Config) {
	logger := logs.NewNullLog()
	sdkClient, err := sdk.NewClient(
		logger,
		sc.ClientPrivateKey,
		sc.ClientPrivateKeyPassword,
		sc.ClientPublicKey,
		sc.Ip,
		sc.Environment,
		sc.ServerPublicKey,
		sc.Address,
		sc.Buffer,
		sc.ZippedMessaging,
		sc.MaxConnections,
		sc.ClientID,
	)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	service := NewSKD(sdkClient)
	checkKeyData(t, sdkClient)

	var (
		respChan = make(chan *structs.KeyManager)
		errChan  = make(chan error)
	)
	for i := 0; i < int(sc.MaxConnections); i++ {
		go func() {
			response, err := service.KeyInfo()
			if err != nil {
				errChan <- err
				return
			}
			respChan <- response
		}()
	}

	select {
	case response := <-respChan:
		if !assert.NotEmpty(t, response.ServerPublicKey) {
			t.FailNow()
		}
	case err := <-errChan:
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	}

}
