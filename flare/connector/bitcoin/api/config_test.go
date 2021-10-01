package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/flarenetwork/coreth/flare/connector/bitcoin/api"
)

func TestWithHost(t *testing.T) {
	host := "host"

	cfg := api.DefaultConfig
	cfg.Host = ""

	api.WithHost(host)(&cfg)
	assert.Equal(t, host, cfg.Host)
}

func TestWithSecure(t *testing.T) {
	secure := true

	cfg := api.DefaultConfig
	cfg.DisableTLS = true

	api.WithSecure(secure)(&cfg)
	assert.Equal(t, !secure, cfg.DisableTLS)
}

func TestWithUser(t *testing.T) {
	user := "user"

	cfg := api.DefaultConfig
	cfg.User = ""

	api.WithUser(user)(&cfg)
	assert.Equal(t, user, cfg.User)
}

func TestWithPassword(t *testing.T) {
	password := "password"

	cfg := api.DefaultConfig
	cfg.Pass = ""

	api.WithPassword(password)(&cfg)
	assert.Equal(t, password, cfg.Pass)
}
