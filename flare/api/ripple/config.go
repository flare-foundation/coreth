package ripple

var DefaultConfig = Config{
	Endpoint: "127.0.0.1:51234",
}

type Config struct {
	Endpoint string
}

type Option func(*Config)

func WithEndpoint(endpoint string) Option {
	return func(cfg *Config) {
		cfg.Endpoint = endpoint
	}
}
