package bitcoin

import (
	"encoding/hex"
	"fmt"

	"github.com/flare-foundation/coreth/flare/connector/bitcoin"
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

// Block retrieves the block with the given hash from the RPC API of the Bitcoin
// node the client is connected to.
func (a *APIClient) Block(hash [32]byte) (*bitcoin.Block, error) {

	res, err := a.client.Call(MethodHeader, hex.EncodeToString(hash[:]), true)
	if err != nil {
		return nil, fmt.Errorf("could not execute call: %w", err)
	}
	if res.Error != nil {
		return nil, fmt.Errorf("could not get response: %w", err)
	}

	var header HeaderResponse
	err = res.GetObject(&header)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	block := bitcoin.Block{
		Hash:          hash,
		Height:        header.Height,
		Confirmations: header.Confirmations,
	}

	return &block, nil
}

// Transaction retrieves the transaction with the given hash and output index
// from the RPC API of the Bitcoin node the client is connected to.
func (a *APIClient) Transaction(hash [32]byte, index uint8) (*bitcoin.Transaction, error) {

	res, err := a.client.Call(MethodTransaction, hex.EncodeToString(hash[:]), false, true)
	if err != nil {
		return nil, fmt.Errorf("could not execute call: %w", err)
	}
	if res.Error != nil {
		return nil, fmt.Errorf("could not get response: %w", err)
	}

	var result TransactionResponse
	err = res.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	// Check whether on output with the given index exists in the transaction
	// with the given hash.
	// NOTE: A Bitcoin transaction can have up 2500-3000 outputs; we thus need
	// to do the comparison using a bigger integer type, otherwise the node
	// would crash if a transaction with more than 255 outputs is processed.
	if uint16(len(result.Decoded.Outputs)) < uint16(index) {
		return nil, bitcoin.ErrOutputNotFound
	}

	// The transaction script has to represent a payment to a public key hash,
	// and a payment to a single address. Otherwise, it's an invalid transaction
	// within the Flare state connector system.
	output := result.Decoded.Outputs[index]
	if output.Key.Type != OutputTypePubKeyHash {
		return nil, bitcoin.ErrInvalidKeyType
	}
	if len(output.Key.Addresses) != NumAddressesRequired {
		return nil, bitcoin.ErrTooManyRecipients
	}

	var block [32]byte
	data, err := hex.DecodeString(result.Block)
	if err != nil {
		return nil, fmt.Errorf("could not decode block hash: %w", err)
	}
	copy(block[:], data)

	recipient := output.Key.Addresses[0]
	amount := uint64(100000000 * output.Value)
	transaction := bitcoin.Transaction{
		Block:         block,
		Hash:          hash,
		Index:         index,
		Recipient:     recipient,
		Amount:        amount,
		Confirmations: result.Confirmations,
	}

	return &transaction, nil
}
