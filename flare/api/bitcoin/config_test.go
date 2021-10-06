package bitcoin_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/flarenetwork/coreth/flare/api/bitcoin"
)

func TestWithHost(t *testing.T) {
	host := "host"

	cfg := bitcoin.DefaultConfig
	cfg.Host = ""

	bitcoin.WithHost(host)(&cfg)
	assert.Equal(t, host, cfg.Host)
}

func TestWithSecure(t *testing.T) {
	secure := true

	cfg := bitcoin.DefaultConfig
	cfg.DisableTLS = true

	bitcoin.WithSecure(secure)(&cfg)
	assert.Equal(t, !secure, cfg.DisableTLS)
}

func TestWithUser(t *testing.T) {
	user := "user"

	cfg := bitcoin.DefaultConfig
	cfg.User = ""

	bitcoin.WithUser(user)(&cfg)
	assert.Equal(t, user, cfg.User)
}

func TestWithPassword(t *testing.T) {
	password := "password"

	cfg := bitcoin.DefaultConfig
	cfg.Pass = ""

	bitcoin.WithPassword(password)(&cfg)
	assert.Equal(t, password, cfg.Pass)
}
