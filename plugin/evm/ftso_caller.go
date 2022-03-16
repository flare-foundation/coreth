// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/core"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/internal/ethapi"
)

type FTSOContract struct {
	address common.Address
	abi     abi.ABI
}

type FTSOCall struct {
	blockchain *core.BlockChain
	hash       common.Hash
	contract   FTSOContract
	err        error
}

type FTSOReturn struct {
	values []interface{}
	err    error
}

type FTSOCaller struct {
	blockchain *core.BlockChain
	hash       common.Hash
}

func NewFTSOCaller(blockchain *core.BlockChain, hash common.Hash) *FTSOCaller {

	f := FTSOCaller{
		blockchain: blockchain,
		hash:       hash,
	}

	return &f
}

func (f *FTSOCaller) OnContract(contract FTSOContract) FTSOCall {
	return FTSOCall{blockchain: f.blockchain, hash: f.hash, contract: contract}
}

func (f FTSOCall) Execute(method string, params ...interface{}) FTSOReturn {

	if f.err != nil {
		return FTSOReturn{err: f.err}
	}

	header := f.blockchain.GetHeaderByHash(f.hash)
	if header == nil {
		return FTSOReturn{err: fmt.Errorf("unknown block (hash: %x)", f.hash)}
	}

	state, err := f.blockchain.StateAt(header.Root)
	if err != nil {
		return FTSOReturn{err: fmt.Errorf("could not get blockchain state (root: %x): %w", header.Root, err)}
	}

	data, err := f.contract.abi.Pack(method, params...)
	if err != nil {
		return FTSOReturn{err: fmt.Errorf("could not pack parameters: %w", err)}
	}

	input := hexutil.Bytes(data)
	args := ethapi.TransactionArgs{To: &f.contract.address, Input: &input}
	msg, err := args.ToMessage(0, nil)
	if err != nil {
		return FTSOReturn{err: fmt.Errorf("could not convert arguments to message: %w", err)}
	}

	txContext := core.NewEVMTxContext(msg)
	blkContext := core.NewEVMBlockContext(header, f.blockchain, nil)
	chainConfig := f.blockchain.Config()
	evm := vm.NewEVM(blkContext, txContext, state, chainConfig, vm.Config{NoBaseFee: true})
	defer evm.Cancel()

	gp := new(core.GasPool).AddGas(math.MaxUint64)
	result, err := core.ApplyMessage(evm, msg, gp)
	if err != nil {
		return FTSOReturn{err: fmt.Errorf("could not apply message: %w", err)}
	}
	if result.Err != nil {
		return FTSOReturn{err: fmt.Errorf("could not execute transaction: %w", err)}
	}

	values, err := f.contract.abi.Unpack(method, result.ReturnData)
	if err != nil {
		return FTSOReturn{err: fmt.Errorf("could not unpack return data: %w", err)}
	}

	return FTSOReturn{values: values}

}

func (f FTSOReturn) Decode(values ...interface{}) error {

	if f.err != nil {
		return f.err
	}

	if len(f.values) != len(values) {
		return fmt.Errorf("invalid number of return values (have: %d, want: %d)", len(f.values), len(values))
	}

	// implement decode reflection logic

	return nil
}
