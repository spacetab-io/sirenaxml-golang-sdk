package sdk

import (
	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk/client"
)

type storage struct {
	c *client.Channel
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
	)

	if err != nil {
		return nil, errors.Wrap(err, "sirena client init error")
	}
	return &storage{c: c}, nil
}

func (s *storage) SendRawRequest(req []byte) ([]byte, error) {
	return s.c.SendMsg(req)
}
