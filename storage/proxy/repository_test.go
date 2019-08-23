package proxy

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
)

var proxyPath string

func tearUp() {
	proxyPath = os.Getenv("PROXY_PATH")
}

func TestMain(m *testing.M) {
	tearUp()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestNewStorage(t *testing.T) {
	logger := logs.NewNullLog()
	proxyStorage := NewStorage(proxyPath, logger, false)
	assert.NotNil(t, proxyStorage.r.GetClient())
}
