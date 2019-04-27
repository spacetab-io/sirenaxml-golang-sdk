package des_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/des"
	"github.com/tmconsulting/sirenaxml-golang-sdk/random"
)

func TestDesEncrypt(t *testing.T) {
	key := []byte(random.String(8))
	origtext := []byte("hello world123563332")

	erytext, err := des.Encrypt(origtext, key)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	destext, err := des.Decrypt(erytext, key)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	assert.Equal(t, string(origtext), string(destext))
}
