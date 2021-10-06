package ripple_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/flarenetwork/coreth/flare/api/ripple"
)

func TestWithEndpoint(t *testing.T) {
	endpoint := "endpoint"

	cfg := ripple.DefaultConfig
	cfg.Endpoint = ""

	ripple.WithEndpoint(endpoint)(&cfg)
	assert.Equal(t, endpoint, cfg.Endpoint)
}
