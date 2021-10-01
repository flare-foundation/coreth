package api

import (
	"github.com/btcsuite/btcd/rpcclient"
)

var DefaultConfig = rpcclient.ConnConfig{
	Host:         "127.0.0.1:8332",
	DisableTLS:   false,
	User:         "",
	Pass:         "",
	HTTPPostMode: true,
}

type Option func(*rpcclient.ConnConfig)

func WithHost(host string) Option {
	return func(cfg *rpcclient.ConnConfig) {
		cfg.Host = host
	}
}

func WithSecure(secure bool) Option {
	return func(cfg *rpcclient.ConnConfig) {
		cfg.DisableTLS = !secure
	}
}
func WithUser(user string) Option {
	return func(cfg *rpcclient.ConnConfig) {
		cfg.User = user
	}
}

func WithPassword(password string) Option {
	return func(cfg *rpcclient.ConnConfig) {
		cfg.Pass = password
	}
}
