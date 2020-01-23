package client

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRespPool_Add(t *testing.T) {
	rp := NewRespPool()

	for i := 0; i < 11; i++ {
		rp.Add(uint32(i))
	}

	assert.Equal(t, 11, len(rp.p))
}

func TestRespPool_SavePacket(t *testing.T) {
	rp := NewRespPool()

	for i := 0; i < 11; i++ {
		rp.Add(uint32(i))
	}

	msgID := uint32(rand.Int31n(10))
	p := &Packet{}

	err := rp.SavePacket(msgID, p)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	P, ok := <-rp.p[msgID]
	if !assert.True(t, ok) {
		t.FailNow()
	}

	assert.Equal(t, p, P)
}

func TestRespPool_GetPacket(t *testing.T) {
	rp := NewRespPool()
	var wg sync.WaitGroup
	for i := uint32(0); i < 11; i++ {
		rp.Add(uint32(i))
	}

	for k := uint32(0); k < 11; k++ {
		wg.Add(1)
		go func() {
			wg.Done()
			p := rp.GetPacket(k)
			assert.Equal(t, k, p.header.MessageID)
		}()
	}
	wg.Wait()
	for j := uint32(0); j < 11; j++ {
		go func() {
			p := &Packet{header: &Header{MessageID: j}}
			err := rp.SavePacket(j, p)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
		}()
	}

}
