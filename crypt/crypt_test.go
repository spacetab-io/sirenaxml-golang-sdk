package crypt

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDataWithServerPubKey(t *testing.T) {
	data := []byte(RandString(8))
	key := []byte(strings.ReplaceAll(os.Getenv("CLIENT_PUBLIC_KEY"), "\\n", "\n"))
	t.Run("success", func(t *testing.T) {
		encKey, err := EncryptDataWithServerPubKey(data, key)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.NotEmpty(t, encKey)
		assert.NotEqual(t, data, encKey)
	})
	t.Run("error on empty key", func(t *testing.T) {
		_, err := EncryptDataWithServerPubKey([]byte("some"), nil)
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
	t.Run("error on unknown key type (not RSA, DSA or ECDSA)", func(t *testing.T) {
		key := []byte(strings.ReplaceAll("1-----BEGIN PUBLIC KEY-----\nbleah\n-----END PUBLIC KEY-----", "\\n", "\n"))
		_, err := EncryptDataWithServerPubKey([]byte("some"), key)
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
}
