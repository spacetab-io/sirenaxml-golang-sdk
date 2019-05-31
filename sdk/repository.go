package sdk

import (
	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/sdk/client"
)

type storage struct {
	c   *client.Channel
	Key []byte
}

func NewClient(sc *sirenaXML.Config, l logs.LogWriter) (*storage, error) {
	c, err := client.NewChannel(sc)
	if err != nil {
		return nil, errors.Wrap(err, "sirena client init error")
	}
	c.SetLogger(l)
	return &storage{c: c, Key: c.Key}, nil
}

func (s *storage) SendRawRequest(req []byte) ([]byte, error) {
	return s.c.SendMsg(req)
}
