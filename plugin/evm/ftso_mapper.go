// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
)

type FTSOMapper struct {
	epochs EpochHandler
	last   uint64
}

func NewFTSOMapper(epochs EpochHandler) *FTSOMapper {

	f := FTSOMapper{
		epochs: epochs,
	}

	return &f
}

func (f *FTSOMapper) ByTimestamp(timestamp uint64) (uint64, error) {

	epoch := f.last
	for {

		start, err := f.epochs.StartTime(epoch)
		if err != nil {
			return 0, fmt.Errorf("could not get start time (epoch: %d): %w", epoch, err)
		}

		end, err := f.epochs.EndTime(epoch)
		if err != nil {
			return 0, fmt.Errorf("could not get end time (epoch: %d): %w", epoch, err)
		}

		if timestamp < start {
			epoch--
			continue
		}

		if timestamp > end {
			epoch++
			continue
		}

		f.last = epoch

		return epoch, nil
	}
}
