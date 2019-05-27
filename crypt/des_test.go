package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDesEncrypt(t *testing.T) {
	key := []byte(RandString(8))
	origtext := []byte("hello world123563332")

	erytext, err := DesEncrypt(origtext, key)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	destext, err := DesDecrypt(erytext, key)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	assert.Equal(t, string(origtext), string(destext))
}
