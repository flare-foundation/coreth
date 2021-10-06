package bitcoin

var DefaultMultiConfig = MultiConfig{
	RequiredMatches: 1,
}

type MultiConfig struct {
	RequiredMatches uint
}

type MultiOption func(*MultiConfig)

func WithRequiredMatches(matches uint) MultiOption {
	return func(cfg *MultiConfig) {
		cfg.RequiredMatches = matches
	}
}
