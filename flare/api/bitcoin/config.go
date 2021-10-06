package bitcoin

import (
	"github.com/btcsuite/btcd/rpcclient"
)

// DefaultConfig represents the default configuration for a Bitcoin API client.
var DefaultConfig = rpcclient.ConnConfig{
	Host:         "127.0.0.1:8332",
	DisableTLS:   false,
	User:         "rpcuser",
	Pass:         "rpcpassword",
	HTTPPostMode: true,
}

// Option represents a configuration option for the Bitcoin API client.
type Option func(*rpcclient.ConnConfig)

// WithHost configures the host of the Bitcoin node's RPC API.
func WithHost(host string) Option {
	return func(cfg *rpcclient.ConnConfig) {
		cfg.Host = host
	}
}

// WithSecure configures whether TLS is used when connecting to the Bitcoin
// node's RPC API.
func WithSecure(secure bool) Option {
	return func(cfg *rpcclient.ConnConfig) {
		cfg.DisableTLS = !secure
	}
}

// WithUser configures the user used when authenticating on the Bitcoin node's
// RPC API.
func WithUser(user string) Option {
	return func(cfg *rpcclient.ConnConfig) {
		cfg.User = user
	}
}

// WithPassword configures the password used when authenticating on the Bitcoin
// node's RPC API.
func WithPassword(password string) Option {
	return func(cfg *rpcclient.ConnConfig) {
		cfg.Pass = password
	}
}
