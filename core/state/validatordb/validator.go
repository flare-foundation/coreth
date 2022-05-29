package validatordb

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
)

// Validator represents an active validator. It is used to store validator
// information.
type Validator struct {
	Providers []common.Address
	NodeID    ids.ShortID
	Weight    uint64
}
