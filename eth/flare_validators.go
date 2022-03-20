// (c) 2019-2020, Ava Labs, Inc. All rights reserved.

package eth

import (
	"github.com/flare-foundation/flare/ids"
)

type FlareValidators interface {
	DefaultValidators() (map[ids.ShortID]uint64, error)
	FTSOValidators(epoch uint64) (map[ids.ShortID]uint64, error)
	ActiveValidators(epoch uint64) (map[ids.ShortID]uint64, error)
}
