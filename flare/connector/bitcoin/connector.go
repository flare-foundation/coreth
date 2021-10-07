package bitcoin

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"

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

	height := binary.BigEndian.Uint64(ret[56:64])
	confirmations := binary.BigEndian.Uint64(ret[88:96])
	hash := flare.Hash(ret[96:128])

	block, err := c.api.Block(hash)
	if err != nil {
		return false, fmt.Errorf("could not get block: %w", err)
	}

	if block.Height != height {
		return false, nil
	}

	if block.Confirmations < confirmations {
		return false, fmt.Errorf("insufficient block confirmations (%d < %d)", block.Confirmations, confirmations)
	}

	return true, nil
}

func (c *Connector) ProvePayment(ret []byte) (bool, error) {

	if len(ret) < 257 {
		return false, nil
	}

	height := binary.BigEndian.Uint64(ret[56:64])
	available := binary.BigEndian.Uint64(ret[88:96])
	fingerprint := flare.Hash(ret[96:128])
	vout := string(ret[192:193])
	hash := flare.Hash(ret[193:257])

	index, err := strconv.ParseUint(vout, 16, 8)
	if err != nil {
		return false, fmt.Errorf("non-hexadecimal vout value (%s)", vout)
	}

	transaction, err := c.api.Transaction(hash, uint8(index))
	if errors.Is(err, ErrOutputNotFound) || errors.Is(err, ErrInvalidKeyType) || errors.Is(err, ErrTooManyRecipients) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("could not get transaction: %w", err)
	}

	block, err := c.api.Block(transaction.Block)
	if err != nil {
		return false, fmt.Errorf("could not get block: %w", err)
	}

	if block.Height >= available {
		return false, fmt.Errorf("unavailable block height (%d >= %d)", block.Height, available)
	}

	if block.Height != height {
		return false, nil
	}

	if transaction.Fingerprint(c.cfg.Currency) != fingerprint {
		return false, nil
	}

	return true, nil
}

func (c *Connector) DisprovePayment(ret []byte) (bool, error) {

	if len(ret) < 257 {
		return false, nil
	}

	height := binary.BigEndian.Uint64(ret[56:64])
	available := binary.BigEndian.Uint64(ret[88:96])
	fingerprint := flare.Hash(ret[96:128])
	vout := string(ret[192:193])
	hash := flare.Hash(ret[193:257])

	index, err := strconv.ParseUint(vout, 16, 8)
	if err != nil {
		return false, fmt.Errorf("non-hexadecimal vout value (%s)", vout)
	}

	transaction, err := c.api.Transaction(hash, uint8(index))
	if errors.Is(err, ErrOutputNotFound) || errors.Is(err, ErrInvalidKeyType) || errors.Is(err, ErrTooManyRecipients) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("could not get transaction: %w", err)
	}

	block, err := c.api.Block(transaction.Block)
	if err != nil {
		return false, fmt.Errorf("could not get block: %w", err)
	}

	if block.Height >= available {
		return false, fmt.Errorf("unavailable block height (%d >= %d)", block.Height, available)
	}

	if block.Height <= height {
		return false, nil
	}

	if transaction.Fingerprint(c.cfg.Currency) != fingerprint {
		return false, nil
	}

	return true, nil
}
