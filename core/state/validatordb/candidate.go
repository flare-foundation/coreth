package validatordb

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
)

type Candidate struct {
	Providers []common.Address
	NodeID    ids.ShortID
	Votepower float64
}
