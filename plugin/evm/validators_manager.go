// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"
)

type Validators interface {
	DefaultValidators(epoch uint64) (map[ids.ShortID]uint64, error)
	FTSOValidators(epoch uint64) (map[ids.ShortID]uint64, error)
	ActiveValidators(epoch uint64) (map[ids.ShortID]uint64, error)
}

type ValidatorsRetriever interface {
	ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error)
}

// ValidatorsManager is responsible for choosing the strategy for building a validator
// set depending on a block hash. It might choose a legacy static validator set, as used
// before the hard fork upgrade, or a dynamic set of validators based on a transition to
// the FTSO validator set.
type ValidatorsManager struct {
	log                  logging.Logger
	defaultValidators    ValidatorsRetriever
	ftsoValidators       ValidatorsRetriever
	activeValidators     ValidatorsRetriever
	transitionValidators ValidatorsRetriever
}

// NewValidatorsManager creates a new manager of validator sets. It uses the given
// blockchain to map block hashes to block headers, the given epoch mapper no map
// block timestamps to FTSO rewards epochs, the given validators as the legacy static
// validator set, and the given retriever to get the validator set based on FTSO
// data providers.
func NewValidatorsManager(log logging.Logger, defaultValidators ValidatorsRetriever, ftsoValidators ValidatorsRetriever, activeValidators ValidatorsRetriever, transitionValidators ValidatorsRetriever) *ValidatorsManager {

	v := ValidatorsManager{
		log:                  log,
		defaultValidators:    defaultValidators,
		ftsoValidators:       ftsoValidators,
		activeValidators:     activeValidators,
		transitionValidators: transitionValidators,
	}

	return &v
}

func (v *ValidatorsManager) DefaultValidators(epoch uint64) (map[ids.ShortID]uint64, error) {
	validators, err := v.defaultValidators.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve default validators: %w", err)
	}
	v.log.Debug("returning default validators")
	return validators, nil
}

func (v *ValidatorsManager) FTSOValidators(epoch uint64) (map[ids.ShortID]uint64, error) {
	validators, err := v.ftsoValidators.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve FTSO validators: %w", err)
	}
	v.log.Debug("returning FTSO validators")
	return validators, nil
}

func (v *ValidatorsManager) ActiveValidators(epoch uint64) (map[ids.ShortID]uint64, error) {
	validators, err := v.activeValidators.ByEpoch(epoch)
	if err == nil {
		v.log.Debug("returning active validators")
		return validators, nil
	}
	if !errors.Is(err, leveldb.ErrNotFound) {
		return nil, fmt.Errorf("could not retrieve active validators: %w", err)
	}
	validators, err = v.transitionValidators.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not transition active validators: %w", err)
	}
	v.log.Debug("returning transitioned validators")
	return validators, nil
}
