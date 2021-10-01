package ripple_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	api "gitlab.com/flarenetwork/coreth/flare/api/ripple"
)

func TestWithEndpoint(t *testing.T) {
	endpoint := "endpoint"

	cfg := api.DefaultConfig
	cfg.Endpoint = ""

	api.WithEndpoint(endpoint)(&cfg)
	assert.Equal(t, endpoint, cfg.Endpoint)
}
