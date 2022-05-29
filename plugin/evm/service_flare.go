package evm

import (
	"context"
	"fmt"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

type FlareAPI struct {
	vm *VM
}

func (api *FlareAPI) Validators(_ context.Context, blockID *ids.ID) (map[string]uint64, error) {

	if blockID == nil {
		latestID, err := api.vm.LastAccepted()
		if err != nil {
			return nil, fmt.Errorf("could not get last accepted block: %w", err)
		}
		blockID = &latestID
	}

	set, err := api.vm.GetValidators(*blockID)
	if err != nil {
		return nil, fmt.Errorf("could not get active validators (block: %x): %w", blockID, err)
	}

	validators := make(map[string]uint64)
	for _, validator := range set.List() {
		validators[validator.ID().PrefixedString(constants.NodeIDPrefix)] = validator.Weight()
	}

	return validators, nil
}
