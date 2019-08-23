package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderFlags_Set(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		h := &Header{}
		h.setFlags(true, false)
		assert.True(t, h.Flags.Has(ZippedResponse|ZippedRequest))
	})
}
