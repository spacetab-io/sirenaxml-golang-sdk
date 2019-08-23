package proxy

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"

	"github.com/sirupsen/logrus"
)

func TestStorage_GetCurrentKeyInfo(t *testing.T) {
	//log := logs.NewNullLog()
	log := logrus.New()
	proxyStorage := NewStorage(proxyPath, log, false)
	reqXML, err := xml.Marshal(&structs.KeyInfoRequest{})
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	resp, sirenaErr, err := proxyStorage.GetCurrentKeyInfo(reqXML)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !assert.Nil(t, sirenaErr) {
		t.FailNow()
	}
	if !assert.NotEmpty(t, resp.Answer) {
		t.FailNow()
	}
	if !assert.NotEmpty(t, resp.Answer.KeyInfo) {
		t.FailNow()
	}
	assert.NotEmpty(t, resp.Answer.KeyInfo.KeyManager.ServerPublicKey)
}

func _TestStorage_GetKeyData(t *testing.T) {
	logger := logs.NewNullLog()
	proxyStorage := NewStorage(proxyPath, logger, false)

	kd, err := proxyStorage.GetKeyData()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	if !assert.NotNil(t, kd) {
		t.FailNow()
	}

	assert.NotEmpty(t, kd.Key)
	assert.NotEmpty(t, kd.ID)
}
