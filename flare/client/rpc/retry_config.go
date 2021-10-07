package rpc

import (
	"time"

	"github.com/cenkalti/backoff/v4"
)

// DefaultRetryConfig represents the default configuration for a retrying
// Bitcoin API client.
var DefaultRetryConfig = backoff.ExponentialBackOff{
	InitialInterval:     1 * time.Second,
	RandomizationFactor: 0,
	Multiplier:          2,
	MaxInterval:         8 * time.Second,
	MaxElapsedTime:      30 * time.Second,
	Stop:                backoff.Stop,
	Clock:               backoff.SystemClock,
}

// RetryOption represents a configuration option for a retrying Bitcoin API
// client.
type RetryOption func(*backoff.ExponentialBackOff)

// WithInitialInterval configures the initial retry interval for a retrying
// Bitcoin API client. It corresponds to interval between the first attempt and
// the second attempt.
func WithInitialInterval(interval time.Duration) RetryOption {
	return func(cfg *backoff.ExponentialBackOff) {
		cfg.InitialInterval = interval
	}
}

// WithMaxInterval configures the maximum retry interval for a retrying Bitcoin
// API client. It corresponds to the maximum wait time between one attempt and
// the next.
func WithMaxInterval(interval time.Duration) RetryOption {
	return func(cfg *backoff.ExponentialBackOff) {
		cfg.MaxInterval = interval
	}
}

// WithMaxElapsed configures the maximum duration that a retrying Bitcoin API
// client will keep retrying. After this duration has elapsed, the client will
// return a failure for the request.
func WithMaxElapsed(duration time.Duration) RetryOption {
	return func(cfg *backoff.ExponentialBackOff) {
		cfg.MaxElapsedTime = duration
	}
}
