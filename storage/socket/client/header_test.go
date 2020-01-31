package client

import (
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/stretchr/testify/assert"

	sirenaXML "github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
)

func TestHeaderFlags_Set(t *testing.T) {
	cfg := &sirenaXML.Config{
		ZippedMessaging: true,
	}

	spew.Dump(cfg)

	h := &Header{}
	h.setFlags(cfg, false)

	spew.Dump(cfg)

	assert.True(t, h.Flags.Has(ZippedResponse|ZippedRequest))
}
