package bitcoin

// DefaultMultiConfig represents the default configuration for a Bitcoin API
// client connected to multiple Bitcoin nodes.
var DefaultMultiConfig = MultiConfig{
	RequiredMatches: 1,
}

// MultiConfig is the configuration of the multi-client, using multiple
// underlying clients to interact with multiple Bitcoin nodes.
type MultiConfig struct {
	RequiredMatches uint
}

// MultiOption is a configuration option for the multi-client.
type MultiOption func(*MultiConfig)

// WithRequiredMatches configures the number of matching responses that need to
// be received by the multi-client to consider a response valid. It needs to be
// set to a number smaller or equal to the number of injected clients.
func WithRequiredMatches(matches uint) MultiOption {
	return func(cfg *MultiConfig) {
		cfg.RequiredMatches = matches
	}
}
