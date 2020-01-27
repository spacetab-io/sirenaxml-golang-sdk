package proxy

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
)

type MockPublisher struct {
}

func (p MockPublisher) PublishLogs(logAttributes map[string]string, request, response []byte) error {
	return nil
}

var proxyPath string

func tearUp() {
	//"https://user:SUrPr5vj@sirena-proxy.dev.tmc24.io/"
	proxyPath = "https://" + os.Getenv("PROXY_CREDS") + "@" + os.Getenv("PROXY_PATH")
}

func TestMain(m *testing.M) {
	tearUp()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestNewStorage(t *testing.T) {
	nl := logs.NewNullLog()
	p := MockPublisher{}

	proxyStorage := NewStorage(p, proxyPath, nl, false)
	assert.NotNil(t, proxyStorage.r.GetClient())
}
