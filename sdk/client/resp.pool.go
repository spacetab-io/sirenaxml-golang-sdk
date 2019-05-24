package client

import (
	"time"

	"github.com/pkg/errors"
)

type RespPool struct {
	p map[uint32]chan *Packet
}

func NewRespPool() *RespPool {
	return &RespPool{p: make(map[uint32]chan *Packet)}
}

func (rp *RespPool) Add(msgID uint32) {
	rp.p[msgID] = make(chan *Packet, 1)
}

func (rp *RespPool) SavePacket(msgID uint32, p *Packet) error {
	select {
	case rp.p[msgID] <- p:
		return nil
	case <-time.After(200 * time.Millisecond):
		return errors.New("save packet timeout")
	}
}

func (rp *RespPool) GetPacket(msgID uint32) *Packet {
	return <-rp.p[msgID]
}
