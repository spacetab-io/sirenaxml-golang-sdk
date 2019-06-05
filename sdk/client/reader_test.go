package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckError(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		incMessage := []byte("all good")
		outMessage, err := checkError(incMessage)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		if !assert.Equal(t, incMessage, outMessage) {
			t.FailNow()
		}

	})
	t.Run("has error", func(t *testing.T) {
		incMessage := []byte(`<?xml version="1.0" encoding="UTF-8"?><sirena><answer pult="ТЕСТ01"><error code="33008">Не найден элемент arrival</error></answer></sirena>`)
		outMessage, err := checkError(incMessage)
		if !assert.Error(t, err) {
			t.FailNow()
		}
		if !assert.Nil(t, outMessage) {
			t.FailNow()
		}

	})
	t.Run("bad error structure", func(t *testing.T) {
		incMessage := []byte(`<?xml version="1.0" encoding="UTF-8"?><error code="33008">Не найден элемент arrival</error>`)
		outMessage, err := checkError(incMessage)
		if !assert.Error(t, err) {
			t.FailNow()
		}
		if !assert.Nil(t, outMessage) {
			t.FailNow()
		}
	})
}
