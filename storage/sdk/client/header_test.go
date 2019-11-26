package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderFlags_Set(t *testing.T) {
	cfg := &sirenaXML.Config{
		ZippedMessaging: true,
	}

	h := &Header{}
	h.setFlags(cfg, false)
	assert.True(t, h.Flags.Has(ZippedResponse|ZippedRequest))
}
