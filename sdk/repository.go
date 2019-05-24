package sdk

import (
	l "github.com/microparts/logs-go"
	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/sdk/client"
)

type storage struct {
	c   *client.Channel
	Key []byte
}

func NewClient(sc *configuration.SirenaConfig, lc *l.Config) (*storage, error) {
	err := logs.Init(lc)
	if err != nil {
		return nil, errors.Wrap(err, "sirena client logging init error")
	}
	c, err := client.NewChannel(sc)
	if err != nil {
		return nil, errors.Wrap(err, "sirena client init error")
	}
	return &storage{c: c, Key: c.Key}, nil
}
