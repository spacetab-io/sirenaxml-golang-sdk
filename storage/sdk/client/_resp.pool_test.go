package client

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/storage/sdk/message"
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
	rm := &message.ReceivedMessage{MessageID: msgID}

	err := rp.SaveMessage(rm)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	m, ok := <-rp.p[msgID]
	if !assert.True(t, ok) {
		t.FailNow()
	}

	assert.Equal(t, rm, m)
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
			defer wg.Done()
			p := rp.GetMessage(k)
			assert.Equal(t, k, p.MessageID)
		}()
	}
	wg.Wait()
	for j := uint32(0); j < 11; j++ {
		go func() {
			p := &message.ReceivedMessage{MessageID: j}
			err := rp.SaveMessage(p)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
		}()
	}

}
