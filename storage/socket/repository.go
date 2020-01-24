package socket

import (
	"bytes"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/socket/client"
)

type storage struct {
	c *client.Channel
	p *http.Client
}

func NewClient(
	l logs.LogWriter,
	clientPrivateKey,
	clientPrivateKeyPassword,
	clientPublicKey,
	ip,
	environment,
	serverPublicKey,
	addr string,
	buffer int,
	zippedMessaging bool,
	maxConnections uint32,
	clientID uint16,
	maxConnectTries int,
) (*storage, error) {

	c, err := client.NewChannel(
		l,
		client.SetClientID(clientID),
		client.SetClientPrivateKey(clientPrivateKey),
		client.SetClientPublicKey(clientPublicKey),
		client.SetClientPrivateKeyPassword(clientPrivateKeyPassword),
		client.SetEnvironment(environment),
		client.SetIp(ip),
		client.SetMaxConnections(maxConnections),
		client.SetSendChannel(buffer),
		client.SetSocket(addr),
		client.SetServerPublicKey(serverPublicKey),
		client.SetZippedMessaging(zippedMessaging),
		client.SetMaxConnectionTries(maxConnectTries),
	)

	st := &storage{c: c}

	if err != nil {
		return nil, errors.Wrap(err, "sirena client init error")
	}
	return st, nil
}

func (s *storage) SendRawRequest(req []byte) ([]byte, error) {
	return s.c.SendMsg(req)
}

// Request sends sirena XML request to sirena proxy
func (s *storage) Request(requestBytes []byte, logAttributes map[string]string) ([]byte, error) {

	request, err := http.NewRequest("POST", "", bytes.NewBuffer(requestBytes))
	if err != nil {
		return nil, err
	}

	response, err := s.p.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}
