package bitcoin

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"

	"gitlab.com/flarenetwork/coreth/flare/connector/bitcoin"
)

const (

	// requiredKeyType is the type of script output for an output to be
	// considered valid for the state connector system. We only support the most
	// basic of Bitcoin transactions, where the output is unlocked with a simple
	// signature of the private key corresponding to the public key hash.
	requiredOutputType = "pubkeyhash"

	// requiredNumAddresses defines the number of signer addresses an output
	// can have to be valid for the state connector system. We only support
	// single payer transactions.
	requiredNumAddresses = 1
)

// Client is a simple Bitcoin API client connecting to the RPC API of a Bitcoin
// node to blocks and transactions. It does not offer any retry functionality.
type Client struct {
	client *rpcclient.Client
}

// NewClient creates a new simple client, connecting to the RPC API of a Bitcoin
// node with the parameters configured by the given options.
func NewClient(options ...Option) (*Client, error) {

	cfg := DefaultConfig
	for _, option := range options {
		option(&cfg)
	}

	client, err := rpcclient.New(&cfg, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create RPC client: %w", err)
	}

	err = client.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not ping RPC server: %w", err)
	}

	c := Client{
		client: client,
	}

	return &c, nil
}

// Block retrieves the block with the given hash from the RPC API of the Bitcoin
// node the client is connected to.
func (c *Client) Block(hash [32]byte) (*bitcoin.Block, error) {

	response, err := c.client.GetBlockHeaderVerbose((*chainhash.Hash)(&hash))
	if err != nil {
		return nil, fmt.Errorf("could not get block: %w", err)
	}

	block := bitcoin.Block{
		Hash:          hash,
		Height:        uint64(response.Height),
		Confirmations: uint64(response.Confirmations),
	}

	return &block, nil
}

// Transaction retrieves the transaction with the given hash and output index
// from the RPC API of the Bitcoin node the client is connected to.
func (c *Client) Transaction(hash [32]byte, index uint8) (*bitcoin.Transaction, error) {

	response, err := c.client.GetRawTransactionVerbose((*chainhash.Hash)(&hash))
	if err != nil {
		return nil, fmt.Errorf("could not get raw transaction: %w", err)
	}

	// Check whether on output with the given index exists in the transaction
	// with the given hash.
	// NOTE: A Bitcoin transaction can have up 2500-3000 outputs; we thus need
	// to do the comparison using a bigger integer type, otherwise the node
	// would crash if a transaction with more than 255 outputs is processed.
	if uint16(len(response.Vout)) < uint16(index) {
		return nil, bitcoin.ErrOutputNotFound
	}

	// The transaction script has to represent a payment to a public key hash,
	// and a payment to a single address. Otherwise, it's an invalid transaction
	// within the Flare state connector system.
	output := response.Vout[index]
	if output.ScriptPubKey.Type != requiredOutputType {
		return nil, bitcoin.ErrInvalidKeyType
	}
	if len(output.ScriptPubKey.Addresses) != requiredNumAddresses {
		return nil, bitcoin.ErrTooManyRecipients
	}

	var block [32]byte
	data, err := hex.DecodeString(response.BlockHash)
	if err != nil {
		return nil, fmt.Errorf("could not decode block hash: %w", err)
	}
	copy(block[:], data)

	recipient := output.ScriptPubKey.Addresses[0]
	amount := uint64(100000000 * output.Value)
	transaction := bitcoin.Transaction{
		Block:         block,
		Hash:          hash,
		Index:         index,
		Recipient:     recipient,
		Amount:        amount,
		Confirmations: response.Confirmations,
	}

	return &transaction, nil
}
