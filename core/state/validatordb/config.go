package validatordb

// DefaultConfig is the default configuration for the validator state database.
var DefaultConfig = Config{
	CacheSize: 10, // 10 MB
}

// Config represents a configuration for the validator state database.
type Config struct {
	CacheSize int // in MB
}

// Option represents an option to be applied against a validator state database
// configuration.
type Option func(*Config)

// WithCacheSize sets the cache size parameter on a validator state database
// configuration. It is given in number of MB, based on how it is used by the
// underlying Ethereum trie library.
func WithCacheSize(size int) Option {
	return func(cfg *Config) {
		cfg.CacheSize = size
	}
}
