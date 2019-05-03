package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetFlags(t *testing.T) {
	t.Run("", func(t *testing.T) {
		p := &NewHeaderParams{MessageIsZipped: true, UseEncrypt: true}
		b := setFlags(p)
		assert.Equal(t, ZippedRequest|EncryptPublic, b.val)
	})
	t.Run("", func(t *testing.T) {
		p := &NewHeaderParams{MessageIsZipped: true}
		b := setFlags(p)
		assert.Equal(t, ZippedRequest, b.val)
	})
	t.Run("", func(t *testing.T) {
		p := &NewHeaderParams{MessageIsZipped: true, UseEncrypt: true, UseSymmetric: true}
		b := setFlags(p)
		assert.Equal(t, ZippedRequest|EncryptPublic|EncryptPublic|EncryptSymmetric, b.val)
	})
}
