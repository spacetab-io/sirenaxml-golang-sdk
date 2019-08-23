package client

import (
	"time"
)

// DelayFunc returns how long to wait before next attempt.
type DelayFunc func(attempt uint) time.Duration

// LinearDelay returns delay value increases linearly depending on the current attempt.
func LinearDelay(initialDelay, maxDelay time.Duration) DelayFunc {
	fn := func(attempt uint) time.Duration {
		delay := time.Duration(attempt) * initialDelay

		return minDuration(delay, maxDelay)
	}

	return fn
}

func minDuration(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}

	return b
}
