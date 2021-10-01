package ripple

import (
	"encoding/hex"
	"fmt"

	"github.com/ybbus/jsonrpc/v2"
	"gitlab.com/flarenetwork/coreth/flare/connector/ripple"
)

const (
	requiredTransactionType = "payment"
)

type Client struct {
	client jsonrpc.RPCClient
}

func NewClient(options ...Option) (*Client, error) {

	cfg := DefaultConfig
	for _, option := range options {
		option(&cfg)
	}

	client := jsonrpc.NewClient(cfg.Endpoint)
	_, err := client.Call("ping", nil)
	if err != nil {
		return nil, fmt.Errorf("could not ping API endpoint: %w", err)
	}

	c := Client{
		client: client,
	}

	return &c, nil
}

func (c *Client) Ledger(index uint32) (*ripple.Ledger, error) {

	request := LedgerRequest{
		LedgerIndex:  index,
		Accounts:     false,
		Full:         false,
		Transactions: false,
		Expand:       false,
		OwnerFunds:   false,
	}

	response, err := c.client.Call(MethodLedger, &request)
	if err != nil {
		return nil, fmt.Errorf("could not execute call: %w", err)
	}
	if response.Error != nil {
		return nil, fmt.Errorf("could not get response: %w", response.Error)
	}

	var result LedgerResponse
	err = response.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("could not decode result: %w", err)
	}

	var hash [32]byte
	data, err := hex.DecodeString(result.LedgerHash)
	if err != nil {
		return nil, fmt.Errorf("could not decode ledger hash: %w", err)
	}
	copy(hash[:], data)

	ledger := ripple.Ledger{
		Hash:      hash,
		Index:     result.LedgerIndex,
		Validated: result.Validated,
	}

	return &ledger, nil
}

func (c *Client) Transaction(hash [32]byte) (*ripple.Transaction, error) {

	request := TransactionRequest{}

	response, err := c.client.Call(MethodTx, &request)
	if err != nil {
		return nil, fmt.Errorf("could not execute call: %w", err)
	}
	if response.Error != nil {
		return nil, fmt.Errorf("could not get response: %w", err)
	}

	var result TransactionResponse
	err = response.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("could not decode result: %w", err)
	}

	if result.TransactionType != requiredTransactionType {
		return nil, fmt.Errorf("invalid transaction type (%s != %s)", result.TransactionType, requiredTransactionType)
	}

	transaction := ripple.Transaction{
		Ledger:      result.LedgerIndex,
		Hash:        hash,
		Validated:   result.Validated,
		Amount:      result.Amount,
		Destination: result.Destination,
		Tag:         result.DestinationTag,
	}

	return &transaction, nil
}
