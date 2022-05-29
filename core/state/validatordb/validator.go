package validatordb

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
)

type Validator struct {
	Providers []common.Address
	NodeID    ids.ShortID
	Weight    uint64
}
