// (c) 2021, Flare Networks Limited. All rights reserved.
//
// This file is a derived work, based on the avalanchego library whose original
// notice appears below. It is distributed under a license compatible with the
// licensing terms of the original code from which it is derived.
// Please see the file LICENSE_AVALABS for licensing terms of the original work.
// Please see the file LICENSE for licensing terms.
//
// (c) 2019-2020, Ava Labs, Inc. All rights reserved.

package evm

import (
	"fmt"
	"math/big"

	"gitlab.com/flarenetwork/coreth/core/state"
	"gitlab.com/flarenetwork/coreth/params"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/crypto"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ethereum/go-ethereum/common"
)

// UnsignedImportTx is an unsigned ImportTx
type UnsignedImportTx struct {
	avax.Metadata
	// ID of the network on which this tx was issued
	NetworkID uint32 `serialize:"true" json:"networkID"`
	// ID of this blockchain.
	BlockchainID ids.ID `serialize:"true" json:"blockchainID"`
	// Which chain to consume the funds from
	SourceChain ids.ID `serialize:"true" json:"sourceChain"`
	// Inputs that consume UTXOs produced on the chain
	ImportedInputs []*avax.TransferableInput `serialize:"true" json:"importedInputs"`
	// Outputs
	Outs []EVMOutput `serialize:"true" json:"outputs"`
}

// InputUTXOs returns the UTXOIDs of the imported funds
func (tx *UnsignedImportTx) InputUTXOs() ids.Set {
	set := ids.Set{}
	for _, in := range tx.ImportedInputs {
		set.Add(in.InputID())
	}
	return set
}

// Verify this transaction is well-formed
func (tx *UnsignedImportTx) Verify(
	avmID ids.ID,
	ctx *snow.Context,
	rules params.Rules,
) error {
	return errWrongChainID
}

func (tx *UnsignedImportTx) Cost() (uint64, error) {
	return 0, fmt.Errorf("exportTx transactions disabled")
}

// Amount of [assetID] burned by this transaction
func (tx *UnsignedImportTx) Burned(assetID ids.ID) (uint64, error) {
	return 0, fmt.Errorf("exportTx transactions disabled")
}

// SemanticVerify this transaction is valid.
func (tx *UnsignedImportTx) SemanticVerify(
	vm *VM,
	stx *Tx,
	parent *Block,
	baseFee *big.Int,
	rules params.Rules,
) error {
	return fmt.Errorf("exportTx transactions disabled")
}

// Accept this transaction and spend imported inputs
// We spend imported UTXOs here rather than in semanticVerify because
// we don't want to remove an imported UTXO in semanticVerify
// only to have the transaction not be Accepted. This would be inconsistent.
// Recall that imported UTXOs are not kept in a versionDB.
func (tx *UnsignedImportTx) Accept(ctx *snow.Context, batch database.Batch) error {
	return fmt.Errorf("exportTx transactions disabled")
}

// newImportTx returns a new ImportTx
func (vm *VM) newImportTx(
	chainID ids.ID, // chain to import from
	to common.Address, // Address of recipient
	baseFee *big.Int, // fee to use post-AP3
	keys []*crypto.PrivateKeySECP256K1R, // Keys to import the funds
) (*Tx, error) {
	return nil, errWrongChainID
}

// EVMStateTransfer performs the state transfer to increase the balances of
// accounts accordingly with the imported EVMOutputs
func (tx *UnsignedImportTx) EVMStateTransfer(ctx *snow.Context, state *state.StateDB) error {
	return errInsufficientFunds
}
