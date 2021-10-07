package multi

var DefaultConfig = Config{
	MatchesRequired: 1,
}

type Config struct {
	MatchesRequired uint
}

type Option func(*Config)

func WithMatchesRequired(matches uint) Option {
	return func(cfg *Config) {
		cfg.MatchesRequired = matches
	}
}
