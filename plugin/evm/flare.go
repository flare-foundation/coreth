// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/flare-foundation/coreth/params"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/coreth/core"
	"github.com/flare-foundation/coreth/core/types"
	"github.com/flare-foundation/coreth/core/vm"
)

type Contract interface {
	Creators(timestamp time.Time) (map[string]uint64, error)
}

type GetCreatorsRequest struct {
	Timestamp time.Time `json:"timestamp"`
}

type GetCreatorsResponse struct {
	Creators map[string]uint64 `json:"creators"`
}

type FlareAPI struct {
	blockchain *core.BlockChain
}

func (f *FlareAPI) GetCreators(_ *http.Request, req *GetCreatorsRequest, res *GetCreatorsResponse) error {
	log.Info("GetCreators called...")
	data := []byte{
		// 4 bytes function code
	}

	msg := types.NewMessage(
		common.Address{},  // from
		&common.Address{}, // to
		0,                 // nonce,
		nil,               // amount
		0,                 // gaslimit
		big.NewInt(5),     // gasprice
		nil,               // gasfeecap
		nil,               // gastipcap
		data,              // data
		nil,               // accesslist
		true,              // isfake
	)

	state, err := f.blockchain.State()
	if err != nil {
		return fmt.Errorf("could not get blockchain state: %w", err)
	}

	tx := core.NewEVMTxContext(msg)
	header := types.Header{
		BaseFee: nil,
		Number:  big.NewInt(1), //todo currently block height is 1. todo: api call needs to give us the block number we care about
		// block number and hash.
		Difficulty: big.NewInt(1),
	}
	//block := core.NewEVMBlockContext(block.Header(), f.blockchain, nil)
	block := core.NewEVMBlockContext(&header, f.blockchain, nil)
	chainConfig := params.ChainConfig{
		ChainID:             big.NewInt(4294967295),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
	}
	evm := vm.NewEVM(block, tx, state, &chainConfig, vm.Config{})
	evmCallValue := big.NewInt(0)
	//evmCallValue = nil
	caller := vm.AccountRef(getCreatorsContractAddress())
	creatorsByte, _, err := evm.Call(caller, getCreatorsContractAddress(), getValidatorsContractFunction4Bytes(), 100000, evmCallValue)
	//caller ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int
	if err != nil {
		return fmt.Errorf("could not get block creators from contract: %w", err)
	}
	log.Info("result: ", "result: ", fmt.Sprintf("%x", creatorsByte))
	creators := make(map[string]uint64)
	err = json.Unmarshal(creatorsByte, &creators)
	log.Info("creatorsByte..: ", "len(creatorsByte)", creatorsByte)
	if err != nil {
		return fmt.Errorf("unmarshalling error while trying to get block creators from contract: %w", err)
	}

	res.Creators = creators

	return nil
}

func getCreatorsContractAddress() common.Address {
	return common.HexToAddress("0x1000000000000000000000000000000000000004")
}

func getCreatorsContractFunction4Bytes() []byte {
	switch {
	default:
		return []byte{0xe6, 0xad, 0xc1, 0xee} //getCreators()
	}
}

func getValidatorsContractFunction4Bytes() []byte {
	switch {
	default:
		return []byte{0xb7, 0xab, 0x4d, 0xb5} //getValidators()
	}
}

//getValidators() b7 ab 4d b5
