package bitcoin

import (
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"

	"gitlab.com/flarenetwork/coreth/flare/connector/bitcoin"
)

// RetryClient is a wrapper around any Bitcoin RPC API client which adds retry
// functionality to its requests.
type RetryClient struct {
	client bitcoin.API
	cfg    backoff.ExponentialBackOff
}

// NewRetryClient wraps an existing new Bitcoin API client and applies the
// retry logic configured by the given options to its block and transaction
// requests.
func NewRetryClient(client bitcoin.API, options ...RetryOption) *RetryClient {

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

// Block retrieves the block with the given hash using the underlying Bitcoin
// API client. It keeps retrying until the configured timeout has been reached.
func (r *RetryClient) Block(hash [32]byte) (*bitcoin.Block, error) {
	for {
		block, err := r.client.Block(hash)
		if err == nil {
			return block, nil
		}
		interval := r.cfg.NextBackOff()
		if interval == backoff.Stop {
			return nil, fmt.Errorf("block request timed out")
		}
		time.Sleep(interval)
	}
}

// Transaction retrieves the transaction with the given hash and output index
// using the underlying Bitcoin API client. It keeps retrying until the
// configured timeout has been reached.
func (r *RetryClient) Transaction(hash [32]byte, index uint8) (*bitcoin.Transaction, error) {
	for {
		transaction, err := r.client.Transaction(hash, index)
		if err == nil {
			return transaction, nil
		}
		interval := r.cfg.NextBackOff()
		if interval == backoff.Stop {
			return nil, fmt.Errorf("transaction request timed out")
		}
		time.Sleep(interval)
	}
}
