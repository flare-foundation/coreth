package validatordb

var DefaultConfig = Config{
	CacheSize: 10, // 10 MB
}

type Config struct {
	CacheSize int
}

type Option func(*Config)

func WithCachSize(size int) Option {
	return func(cfg *Config) {
		cfg.CacheSize = size
	}
}
