// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
)

type Epochs interface {
	Details(epoch uint64) (EpochDetails, error)
}

type EpochDetails struct {
	PowerHeight uint64
	StartHeight uint64
	StartTime   uint64
	EndTime     uint64
}

type EpochsManager struct {
	epochs Epochs
	last   uint64
}

func NewEpochsManager(epochs Epochs) *EpochsManager {

	e := EpochsManager{
		epochs: epochs,
		last:   0,
	}

	return &e
}

func (e *EpochsManager) ByTimestamp(timestamp uint64) (uint64, error) {

	epoch := e.last
	for {

		details, err := e.epochs.Details(epoch)
		if err != nil {
			return 0, fmt.Errorf("could not get epoch details for mapping: %w", err)
		}

		if timestamp < details.StartTime {
			epoch--
			continue
		}

		if timestamp > details.EndTime {
			epoch++
			continue
		}

		e.last = epoch

		return epoch, nil
	}
}
