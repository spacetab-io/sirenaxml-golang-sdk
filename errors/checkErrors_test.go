package errors

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestCheckError(t *testing.T) {
	err := errors.New("some error")
	type errDesc struct {
		body     []byte
		code     int
		cryptErr int
		mess     string
	}
	respErrors := map[string]errDesc{
		"key error": {
			body:     []byte(`<?xml version="1.0" encoding="UTF-8"?><sirena><answer><error code="-1" crypt_error="5">5 Unknown symmetric key</error></answer></sirena>`),
			code:     -1,
			cryptErr: 5,
			mess:     "5 Unknown symmetric key",
		},
	}
	t.Run("success error find", func(t *testing.T) {
		hasErr, respErr, err := CheckErrors(nil, err)

		assert.True(t, hasErr)
		assert.Error(t, err)
		assert.Nil(t, respErr)
	})
	for name, errXml := range respErrors {
		t.Run(fmt.Sprintf("success resp error find: %s", name), func(t *testing.T) {
			hasErr, respErr, err := CheckErrors(errXml.body, nil)

			assert.True(t, hasErr)
			assert.Nil(t, err)

			assert.Equal(t, errXml.mess, respErr.Message)
			assert.Equal(t, errXml.cryptErr, respErr.CryptError)
			assert.Equal(t, errXml.code, respErr.Code)
		})
	}
}
