package validators

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
)

type Entry struct {
	Provider  common.Address
	NodeID    ids.ShortID
	Votepower float64
}
