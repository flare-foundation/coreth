package ripple

import (
	"fmt"
)

type Connector struct {
	api API
	cfg Config
}

func NewConnector(api API, options ...Option) *Connector {

	cfg := DefaultConfig
	for _, option := range options {
		option(&cfg)
	}

	c := Connector{
		api: api,
		cfg: cfg,
	}

	return &c
}

func (c *Connector) ProveDataAvailabilityPeriodFinality(ret []byte) (bool, error) {

	// TODO
	var height uint32
	var hash [32]byte

	ledger, err := c.api.Ledger(height)
	if err != nil {
		return false, fmt.Errorf("could not get ledger: %w", err)
	}

	if ledger.Hash != hash {
		return false, fmt.Errorf("invalid ledger hash (%x != %x)", ledger.Hash, hash)
	}

	return true, nil
}

func (c *Connector) ProvePaymentFinality(ret []byte) (bool, error) {

	// TODO
	var hash [32]byte
	var fingerprint [32]byte

	transaction, err := c.api.Transaction(hash)
	if err != nil {
		return false, fmt.Errorf("could not get transaction: %w", err)
	}

	if transaction.Fingerprint(c.cfg.Currency) != fingerprint {
		return false, fmt.Errorf("invalid transaction fingerprint (%x != %x)", transaction.Fingerprint(c.cfg.Currency), fingerprint)
	}

	return true, nil
}

func (c *Connector) DisprovePaymentFinality(ret []byte) (bool, error) {

	return true, nil
}
