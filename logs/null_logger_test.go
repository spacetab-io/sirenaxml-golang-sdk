package logs

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNullLog(t *testing.T) {
	logger := NewNullLog()
	assert.Empty(t, captureOutput(func() {
		logger.Debug("debug")
	}))
	assert.Empty(t, captureOutput(func() {
		logger.Debugf("debug")
	}))
	assert.Empty(t, captureOutput(func() {
		logger.Info("info")
	}))
	assert.Empty(t, captureOutput(func() {
		logger.Infof("info")
	}))
	assert.Empty(t, captureOutput(func() {
		logger.Warning("warning")
	}))
	assert.Empty(t, captureOutput(func() {
		logger.Warningf("warning")
	}))
	assert.Empty(t, captureOutput(func() {
		logger.Error("error")
	}))
	assert.Empty(t, captureOutput(func() {
		logger.Errorf("error")
	}))
	assert.Empty(t, captureOutput(func() {
		assert.Panics(t, func() { logger.Fatal("fatal") })
	}))
	assert.Empty(t, captureOutput(func() {
		assert.Panics(t, func() { logger.Fatalf("fatal") })
	}))
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
