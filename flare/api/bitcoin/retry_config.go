package bitcoin

import (
	"time"

	"github.com/cenkalti/backoff/v4"
)

var DefaultRetryConfig = backoff.ExponentialBackOff{
	InitialInterval:     1 * time.Second,
	RandomizationFactor: 0,
	Multiplier:          2,
	MaxInterval:         8 * time.Second,
	MaxElapsedTime:      30 * time.Second,
	Stop:                backoff.Stop,
	Clock:               backoff.SystemClock,
}

type RetryOption func(*backoff.ExponentialBackOff)

func WithInitialInterval(interval time.Duration) RetryOption {
	return func(cfg *backoff.ExponentialBackOff) {
		cfg.InitialInterval = interval
	}
}

func WithMaxInterval(interval time.Duration) RetryOption {
	return func(cfg *backoff.ExponentialBackOff) {
		cfg.MaxInterval = interval
	}
}

func WithMaxElapsed(duration time.Duration) RetryOption {
	return func(cfg *backoff.ExponentialBackOff) {
		cfg.MaxElapsedTime = duration
	}
}
