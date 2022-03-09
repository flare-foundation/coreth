// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/core"
	evm "github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/internal/ethapi"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validators"
)

type FTSO struct {
	blockchain *core.BlockChain
	submitter  abi.ABI
	registry   abi.ABI
	manager    abi.ABI
	asset      abi.ABI
}

type ValidatorSnapshot struct {
	start      uint64
	end        uint64
	validators validators.Set
}

func NewFTSO(vm *VM, address common.Address) (*FTSO, error) {

	submitter, err := abi.JSON(strings.NewReader(abiPriceSubmitter))
	if err != nil {
		return nil, fmt.Errorf("could not decode price submitter ABI: %w", err)
	}
	registry, err := abi.JSON(strings.NewReader(abiFTSORegistry))
	if err != nil {
		return nil, fmt.Errorf("could not decode FTSO registry ABI: %w", err)
	}
	manager, err := abi.JSON(strings.NewReader(abiFTSOManager))
	if err != nil {
		return nil, fmt.Errorf("could not decode FTSO manager ABI: %w", err)
	}
	asset, err := abi.JSON(strings.NewReader(abiFTSOAsset))
	if err != nil {
		return nil, fmt.Errorf("could not decode FTSO asset ABI: %w", err)
	}

	f := FTSO{
		blockchain: vm.chain.BlockChain(),
		submitter:  submitter,
		registry:   registry,
		manager:    manager,
		asset:      asset,
	}

	return &f, nil
}

func (f *FTSO) Epoch(time uint64) (uint64, error) {

	return 0, nil
}

func (f *FTSO) Validators(epoch uint64) ([]ids.ShortID, error) {
	return nil, nil
}

func (f *FTSO) Rewards(validatorID ids.ShortID, epoch uint64) (uint64, error) {
	return 0, nil
}

func (f *FTSO) call(hash common.Hash, abi abi.ABI, method string, params ...interface{}) ([]interface{}, error) {

	data, err := abi.Pack(method, params...)
	if err != nil {
		return nil, fmt.Errorf("could not pack call data: %w", err)
	}

	header := f.blockchain.GetHeaderByHash(hash)
	if header == nil {
		return nil, fmt.Errorf("block hash unknown")
	}

	state, err := f.blockchain.StateAt(header.Root)
	if err != nil {
		return nil, fmt.Errorf("could not get blockchain state: %w", err)
	}

	input := hexutil.Bytes(data)
	args := ethapi.TransactionArgs{Input: &input}
	msg, err := args.ToMessage(0, nil)
	if err != nil {
		return nil, fmt.Errorf("could not convert arguments to message: %w", err)
	}

	vmConfig := f.blockchain.GetVMConfig()
	chainConfig := f.blockchain.Config()
	txContext := core.NewEVMTxContext(msg)
	blkContext := core.NewEVMBlockContext(header, f.blockchain, nil)
	evm := evm.NewEVM(blkContext, txContext, state, chainConfig, *vmConfig)
	defer evm.Cancel()

	gp := new(core.GasPool).AddGas(math.MaxUint64)
	result, err := core.ApplyMessage(evm, msg, gp)
	if err != nil {
		return nil, fmt.Errorf("could not apply message: %w", err)
	}
	if result.Err != nil {
		return nil, fmt.Errorf("could not execute transaction: %w", err)
	}

	values, err := abi.Unpack(method, result.ReturnData)
	if err != nil {
		return nil, fmt.Errorf("could not unpack return data: %w", err)
	}

	return values, nil
}
