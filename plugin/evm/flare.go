// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/coreth/core"
	"github.com/flare-foundation/coreth/core/types"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/flare/vms/proposervm/block"
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

	data := []byte{
		// 4 bytes function code
	}

	msg := types.NewMessage(
		common.Address{},  // from
		&common.Address{}, // to
		0,                 // nonce,
		nil,               // amount
		0,                 // gaslimit
		nil,               // gasprice
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
	block := core.NewEVMBlockContext(block.Header(), f.blockchain, nil)
	evm := vm.NewEVM(block, tx, state, nil, nil)

	creators, err := evm.Call()
	if err != nil {
		return fmt.Errorf("could not get block creators from contract: %w", err)
	}

	res.Creators = creators

	return nil
}
