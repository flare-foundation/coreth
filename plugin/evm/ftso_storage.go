// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"github.com/flare-foundation/coreth/core/state"
	"github.com/flare-foundation/flare/ids"
)

type FTSOStorage struct {
	state *state.StateDB
}

func (f *FTSOStorage) Save(epoch uint64, validators map[ids.ShortID]uint64) error {
	return nil
}

func (f *FTSOStorage) Load() (uint64, map[ids.ShortID]uint64, error) {
	return 0, nil, nil
}
