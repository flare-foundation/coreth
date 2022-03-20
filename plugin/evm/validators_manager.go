// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"

	"github.com/flare-foundation/flare/ids"
)

type ValidatorRetriever interface {
	ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error)
}

// ValidatorsManager is responsible for choosing the strategy for building a validator
// set depending on a block hash. It might choose a legacy static validator set, as used
// before the hard fork upgrade, or a dynamic set of validators based on a transition to
// the FTSO validator set.
type ValidatorsManager struct {
	defaultValidators map[ids.ShortID]uint64
	ftsoValidators    ValidatorRetriever
	activeValidators  ValidatorRetriever
}

// NewValidatorsManager creates a new manager of validator sets. It uses the given
// blockchain to map block hashes to block headers, the given epoch mapper no map
// block timestamps to FTSO rewards epochs, the given validators as the legacy static
// validator set, and the given retriever to get the validator set based on FTSO
// data providers.
func NewValidatorsManager(defaultValidators map[ids.ShortID]uint64, ftsoValidators ValidatorRetriever, activeValidators ValidatorRetriever, opts ...CacheOption) *ValidatorsManager {

	v := ValidatorsManager{
		defaultValidators: defaultValidators,
		ftsoValidators:    ftsoValidators,
		activeValidators:  activeValidators,
	}

	return &v
}

func (v *ValidatorsManager) DefaultValidators() (map[ids.ShortID]uint64, error) {
	return v.defaultValidators, nil
}

func (v *ValidatorsManager) FTSOValidators(epoch uint64) (map[ids.ShortID]uint64, error) {
	validators, err := v.ftsoValidators.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve FTSO validators by epoch: %w", err)
	}
	return validators, nil
}

func (v *ValidatorsManager) ActiveValidators(epoch uint64) (map[ids.ShortID]uint64, error) {
	validators, err := v.activeValidators.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve active validators by epoch: %w", err)
	}
	return validators, nil
}
