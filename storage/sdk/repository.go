package sdk

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk/client"
)

type storage struct {
	ctx context.Context
	c   *client.Client
}

func Client(sc *sirenaXML.Config, l logs.LogWriter) (*storage, error) {
	err := sc.PrepareKeys()
	if err != nil {
		return nil, err
	}
	addr, err := sc.GetAddr()
	if err != nil {
		return nil, err
	}
	// @TODO make timeout configurable?
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	c, err := client.New(
		client.SetKeys(sc.GetKeys()),
		client.SetAddr(addr),
		client.SetLogger(l),
		client.SetUseZip(sc.ZippedMessaging),
		client.SetClientID(sc.ClientID),
		client.SetMaxConnections(sc.MaxConnections),
	)
	if err != nil {
		return nil, errors.Wrap(err, "sirena client init error")
	}

	errCh := make(chan error)
	go func() {
		errCh <- c.Connect(ctx)
	}()
	go func() {
		for {
			select {
			case <-errCh:
				ctx.Done()
			}
		}
	}()
	err = c.WaitKeySign(ctx)

	return &storage{c: c, ctx: ctx}, err
}

func (s *storage) SendRawRequest(req []byte) ([]byte, error) {
	return s.c.SendMsg(s.ctx, req)
}
