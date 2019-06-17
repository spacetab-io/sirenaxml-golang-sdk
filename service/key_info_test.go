package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk"
)

func TestService_KeyInfo(t *testing.T) {
	logger := logs.NewNullLog()
	sdkClient, err := sdk.NewClient(&sc, logger)
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
