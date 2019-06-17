package sdk

import (
	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk/client"
)

type storage struct {
	c *client.Channel
}

func NewClient(sc *sirenaXML.Config, l logs.LogWriter) (*storage, error) {
	c, err := client.NewChannel(sc, l)
	if err != nil {
		return nil, errors.Wrap(err, "sirena client init error")
	}
	return &storage{c: c}, nil
}

func (s *storage) SendRawRequest(req []byte) ([]byte, error) {
	return s.c.SendMsg(req)
}
