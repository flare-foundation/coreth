package ripple

import (
	"encoding/hex"
	"fmt"

	"gitlab.com/flarenetwork/coreth/flare/connector/ripple"
)

type APIClient struct {
	client RPCClient
}

func NewAPIClient(client RPCClient) *APIClient {

	a := APIClient{
		client: client,
	}

	return &a
}

func (a *APIClient) Ledger(index uint32) (*ripple.Ledger, error) {

	req := LedgerRequest{
		LedgerIndex:  index,
		Accounts:     false,
		Full:         false,
		Transactions: false,
		Expand:       false,
		OwnerFunds:   false,
	}

	res, err := a.client.Call(MethodLedger, &req)
	if err != nil {
		return nil, fmt.Errorf("could not execute call: %w", err)
	}
	if res.Error != nil {
		return nil, fmt.Errorf("could not get response: %w", res.Error)
	}

	var result LedgerResponse
	err = res.GetObject(&result)
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
		Height:    result.LedgerIndex,
		Validated: result.Validated,
	}

	return &ledger, nil
}

func (a *APIClient) Transaction(hash [32]byte) (*ripple.Transaction, error) {

	req := TransactionRequest{}

	res, err := a.client.Call(MethodTx, &req)
	if err != nil {
		return nil, fmt.Errorf("could not execute call: %w", err)
	}
	if res.Error != nil {
		return nil, fmt.Errorf("could not get response: %w", err)
	}

	var result TransactionResponse
	err = res.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("could not decode result: %w", err)
	}

	if result.TransactionType != TransactionTypePayment {
		return nil, fmt.Errorf("invalid transaction type (%s != %s)", result.TransactionType, TransactionTypePayment)
	}

	transaction := ripple.Transaction{
		Height:      result.LedgerIndex,
		Hash:        hash,
		Validated:   result.Validated,
		Amount:      result.Amount,
		Destination: result.Destination,
		Tag:         result.DestinationTag,
	}

	return &transaction, nil
}
