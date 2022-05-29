package validatordb

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
)

// Candidate represents a candidate to become an active validator. It is used to
// store transient information on epoch switchover that needs to be available
// when validator weights are calculated.
type Candidate struct {
	Providers []common.Address
	NodeID    ids.ShortID
	Votepower float64
}
