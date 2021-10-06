package bitcoin

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"

	"gitlab.com/flarenetwork/coreth/flare/connector/bitcoin"
)

const (
	requiredKeyType      = "pubkeyhash"
	requiredNumAddresses = 1
)

type Client struct {
	client *rpcclient.Client
}

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

func (c *Client) Transaction(hash [32]byte, index uint8) (*bitcoin.Transaction, error) {

	response, err := c.client.GetRawTransactionVerbose((*chainhash.Hash)(&hash))
	if err != nil {
		return nil, fmt.Errorf("could not get raw transaction: %w", err)
	}

	if uint8(len(response.Vout)) < index {
		return nil, bitcoin.ErrOutputNotFound
	}

	output := response.Vout[index]
	if output.ScriptPubKey.Type != requiredKeyType {
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
