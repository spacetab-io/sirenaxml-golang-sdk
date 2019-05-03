package client

import (
	"math/rand"
	"time"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
)

type msgIDsPool struct {
	c chan uint32
}

const MaxMsgs = 100

func makeMsgPool() *msgIDsPool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := &msgIDsPool{c: make(chan uint32, MaxMsgs)}
	// Fill the msgIDsPool: create objects in advance
	for i := 0; i < cap(p.c); i++ {
		p.c <- r.Uint32()
	}

	logs.Log.Debugf("pool cap: %d, len: %d", cap(p.c), len(p.c))

	return p
}

func (p *msgIDsPool) getMsgIDFromPool() uint32 {
	re := <-p.c
	logs.Log.Debugf("pool cap: %d, len: %d", cap(p.c), len(p.c))
	return re
}

func (p *msgIDsPool) returnMsgIDToPool(id uint32) {
	p.c <- id
	logs.Log.Debugf("pool cap: %d, len: %d", cap(p.c), len(p.c))
}
