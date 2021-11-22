package bitcoin

import (
	"github.com/flare-foundation/coreth/flare"
)

var DefaultConfig = Config{
	Currency: flare.CurrencyBitcoin,
}

type Config struct {
	Currency string
}

type Option func(*Config)

func WithCurrency(currency string) Option {
	return func(cfg *Config) {
		cfg.Currency = currency
	}
}
