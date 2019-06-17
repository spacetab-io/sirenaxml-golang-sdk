package proxy

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/structs"
)

func TestStorage_GetCurrentKeyInfo(t *testing.T) {
	nl := logs.NewNullLog()
	proxyStorage := NewStorage(proxyPath, nl, false)
	reqXML, err := xml.Marshal(&structs.KeyInfoRequest{})
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	resp, err := proxyStorage.GetCurrentKeyInfo(reqXML)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !assert.NotNil(t, resp) {
		t.FailNow()
	}
	assert.NotEmpty(t, resp.Answer.KeyInfo.KeyManager.ServerPubliKey)
}

func _TestStorage_GetKeyData(t *testing.T) {
	nl := logs.NewNullLog()
	proxyStorage := NewStorage(proxyPath, nl, false)

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
