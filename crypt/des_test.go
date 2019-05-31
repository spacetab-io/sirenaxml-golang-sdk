package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDesEncrypt(t *testing.T) {
	key := []byte(RandString(8))
	origtext := []byte("hello world123563332")
	t.Run("success", func(t *testing.T) {
		erytext, err := DesEncrypt(origtext, key)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		destext, err := DesDecrypt(erytext, key)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.Equal(t, string(origtext), string(destext))
	})

	t.Run("error empty key", func(t *testing.T) {
		_, err := DesEncrypt(origtext, nil)
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
	t.Run("error orig key", func(t *testing.T) {
		_, err := DesEncrypt(nil, key)
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
}
