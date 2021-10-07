package rpc

import (
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ybbus/jsonrpc/v2"
)

// RetryClient is a wrapper for RPC clients that adds retry functionality to
// them.
type RetryClient struct {
	client jsonrpc.RPCClient
	cfg    backoff.ExponentialBackOff
}

// NewRetryClient creates a new retry wrapper around the given RPC client,
// configured with the given retry options.
func NewRetryClient(client jsonrpc.RPCClient, options ...RetryOption) *RetryClient {

	cfg := DefaultRetryConfig
	for _, option := range options {
		option(&cfg)
	}

	r := RetryClient{
		client: client,
		cfg:    cfg,
	}

	return &r
}

// Call executes an RPC call on the underlying RPC client and keeps retrying
// according to the configured retry options.
func (r *RetryClient) Call(method string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	for {
		res, err := r.client.Call(method, params...)
		if err == nil {
			return res, nil
		}
		interval := r.cfg.NextBackOff()
		if interval == backoff.Stop {
			return nil, fmt.Errorf("RPC request retry limit reached")
		}
		time.Sleep(interval)
	}
}
