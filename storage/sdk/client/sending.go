package client

import (
	"context"

	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk/message"
)

type sending struct {
	errCh chan error
	msg   message.Message
}

func (q sending) done(ctx context.Context, err error) {
	select {
	case <-ctx.Done():
		return
	case q.errCh <- err:
	}
}

func (q sending) wait(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-q.errCh:
		return err
	}
}
