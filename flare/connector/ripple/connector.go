package ripple

import (
	"encoding/binary"
	"fmt"

	"gitlab.com/flarenetwork/coreth/flare"
)

type Connector struct {
	api APIClient
	cfg Config
}

func NewConnector(api APIClient, options ...Option) *Connector {

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

func (c *Connector) ProveAvailability(ret []byte) (bool, error) {

	if len(ret) < 128 {
		return false, nil
	}

	height := uint32(binary.BigEndian.Uint64(ret[56:64]))
	fingerprint := flare.Hash(ret[96:128])

	ledger, err := c.api.Ledger(height)
	if err != nil {
		return false, fmt.Errorf("could not get ledger: %w", err)
	}

	if ledger.Height < height {
		return false, nil
	}

	if ledger.Fingerprint() != fingerprint {
		return false, nil
	}

	return true, nil
}

func (c *Connector) ProvePayment(ret []byte) (bool, error) {

	if len(ret) < 224 {
		return false, nil
	}

	height := uint32(binary.BigEndian.Uint64(ret[56:64]))
	available := uint32(binary.BigEndian.Uint64(ret[88:96]))
	fingerprint := flare.Hash(ret[96:128])
	hash := flare.Hash(ret[192:224])

	transaction, err := c.api.Transaction(hash)
	if err != nil {
		return false, fmt.Errorf("could not get transaction: %w", err)
	}

	if transaction.Height > available {
		return false, fmt.Errorf("unavailable block height")
	}

	if transaction.Height != height {
		return false, nil
	}

	if transaction.Fingerprint(c.cfg.Currency) != fingerprint {
		return false, nil
	}

	return true, nil
}

func (c *Connector) DisprovePayment(ret []byte) (bool, error) {

	if len(ret) < 224 {
		return false, nil
	}

	height := uint32(binary.BigEndian.Uint64(ret[56:64]))
	available := uint32(binary.BigEndian.Uint64(ret[88:96]))
	fingerprint := flare.Hash(ret[96:128])
	hash := flare.Hash(ret[192:224])

	transaction, err := c.api.Transaction(hash)
	if err != nil {
		return false, fmt.Errorf("could not get transaction: %w", err)
	}

	if transaction.Height > available {
		return false, fmt.Errorf("unavailable block height")
	}

	if transaction.Height <= height {
		return false, nil
	}

	if transaction.Fingerprint(c.cfg.Currency) != fingerprint {
		return false, nil
	}

	return true, nil
}
