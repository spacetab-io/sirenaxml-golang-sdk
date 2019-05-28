package client

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
)

func TestHeaderFlags_Set(t *testing.T) {
	cfg := &configuration.SirenaConfig{
		ZippedMessaging: true,
	}

	h := &Header{}
	h.setFlags(cfg, false)
	assert.True(t, h.Flags.Has(ZippedResponse|ZippedRequest))
}
