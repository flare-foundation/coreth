package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/flarenetwork/coreth/flare/connector/ripple/api"
)

func TestWithEndpoint(t *testing.T) {
	endpoint := "endpoint"

	cfg := api.DefaultConfig
	cfg.Endpoint = ""

	api.WithEndpoint(endpoint)(&cfg)
	assert.Equal(t, endpoint, cfg.Endpoint)
}
