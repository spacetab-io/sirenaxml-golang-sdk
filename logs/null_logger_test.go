package logs

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNullLog(t *testing.T) {
	nl := NewNullLog()
	assert.Empty(t, captureOutput(func() {
		nl.Debug("debug")
	}))
	assert.Empty(t, captureOutput(func() {
		nl.Info("info")
	}))
	assert.Empty(t, captureOutput(func() {
		nl.Warning("warning")
	}))
	assert.Empty(t, captureOutput(func() {
		nl.Error("error")
	}))
	assert.Empty(t, captureOutput(func() {
		nl.Fatal("fatal")
	}))
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
