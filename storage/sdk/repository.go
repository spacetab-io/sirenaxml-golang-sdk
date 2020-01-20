package sdk

import (
	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk/client"
)

type storage struct {
	c *client.Channel
}

func NewClient(l logs.LogWriter, opts ...client.Option) (*storage, error) {

	c, err := client.NewChannel(l, opts)
	if err != nil {
		return nil, errors.Wrap(err, "sirena client init error")
	}
	return &storage{c: c}, nil
}

func (s *storage) SendRawRequest(req []byte) ([]byte, error) {
	return s.c.SendMsg(req)
}
