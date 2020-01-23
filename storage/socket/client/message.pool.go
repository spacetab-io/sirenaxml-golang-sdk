package client

import (
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

type MsgPool struct {
	msgsIDs chan uint32
}

func NewMsgPool(rp *RespPool, size uint32) (*MsgPool, error) {
	if size == 0 {
		return nil, errors.New("empty message pool is not allowed")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := &MsgPool{msgsIDs: make(chan uint32, size)}
	for i := 0; i < cap(p.msgsIDs); i++ {
		msgID := r.Uint32()
		rp.Add(msgID)
		p.msgsIDs <- msgID
	}

	return p, nil
}

func (p *MsgPool) GetPool() chan uint32 {
	return p.msgsIDs
}

func (p *MsgPool) GetMsgID() uint32 {
	return <-p.msgsIDs
}

func (p *MsgPool) ReturnMsgIDToPool(id uint32) {
	p.msgsIDs <- id
}
