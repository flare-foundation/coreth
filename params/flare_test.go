// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package params

import (
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

func TestNewValidatorsDefault(t *testing.T) {

	customNodeIDs := []string{
		"NodeID-3M9KVT6ixi4gVMisbm5TnPXYXgFN5LHuv",
		"NodeID-NnX4fajAmyvpL9RLfheNdc47FKKDuQW8i",
	}

	setValid := func(t *testing.T) {
		err := os.Setenv("CUSTOM_VALIDATORS", strings.Join(customNodeIDs, ","))
		require.NoError(t, err)
	}

	setInvalid := func(t *testing.T) {
		err := os.Setenv("CUSTOM_VALIDATORS", "NodeD-3M9KVT6ixi4gVMisbm5TnPXYXgFN5LHuv")
		require.NoError(t, err)
	}

	tests := []struct {
		name    string
		setEnv  func(t *testing.T)
		nodeIDs []string
		chainID *big.Int
		weight  uint64
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "testing default validators",
			chainID: TestingChainID,
			nodeIDs: testingNodeIDs,
			weight:  testingValidatorWeight,
			wantErr: require.NoError,
		},
		{
			name:    "coston default validators",
			chainID: CostonChainID,
			nodeIDs: costonNodeIDs,
			weight:  costonValidatorWeight,
			wantErr: require.NoError,
		},
		{
			name:    "songbird default validators",
			chainID: SongbirdChainID,
			nodeIDs: songbirdNodeIDs,
			weight:  songbirdValidatorWeight,
			wantErr: require.NoError,
		},
		{
			name:    "flare default validators",
			chainID: FlareChainID,
			nodeIDs: flareNodeIDs,
			weight:  flareValidatorWeight,
			wantErr: require.Error, // flare main network not active yet
		},
		{
			name:    "valid custom default validators",
			setEnv:  setValid,
			chainID: big.NewInt(1337),
			nodeIDs: customNodeIDs,
			weight:  customValidatorWeight,
			wantErr: require.NoError,
		},
		{
			name:    "invalid custom default validators",
			setEnv:  setInvalid,
			chainID: big.NewInt(1337),
			nodeIDs: customNodeIDs,
			weight:  customValidatorWeight,
			wantErr: require.Error,
		},
		{
			name:    "missing custom default validators",
			chainID: big.NewInt(1337),
			nodeIDs: customNodeIDs,
			weight:  customValidatorWeight,
			wantErr: require.Error,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {

			err := os.Unsetenv("CUSTOM_VALIDATORS")
			require.NoError(t, err)

			if test.setEnv != nil {
				test.setEnv(t)
			}

			got, err := NewFlareConfig(test.chainID).DefaultValidators()
			test.wantErr(t, err)

			if err != nil {
				return
			}

			validatorIDs := make([]ids.ShortID, 0, len(test.nodeIDs))
			for _, nodeID := range test.nodeIDs {
				validatorID, err := ids.ShortFromPrefixedString(nodeID, constants.NodeIDPrefix)
				require.NoError(t, err)
				validatorIDs = append(validatorIDs, validatorID)
			}

			for _, validatorID := range validatorIDs {
				assert.Contains(t, got.validators, validatorID)
			}

			for _, weight := range got.validators {
				assert.Equal(t, test.weight, weight)
			}
		})
	}
}

func TestValidatorsDefault_ByEpoch(t *testing.T) {

	validatorIDs := []ids.ShortID{
		{1},
		{2},
		{3},
	}

	validators := map[ids.ShortID]uint64{
		validatorIDs[0]: 1,
		validatorIDs[1]: 2,
		validatorIDs[2]: 3,
	}

	tests := []struct {
		name         string
		steps        []Step
		epoch        uint64
		validatorIDs []ids.ShortID
	}{
		{
			name:         "no steps zero epoch",
			steps:        []Step{},
			epoch:        0,
			validatorIDs: validatorIDs,
		},
		{
			name:         "no steps non-zero epoch",
			steps:        []Step{},
			epoch:        10,
			validatorIDs: validatorIDs,
		},
		{
			name:  "one step epoch before",
			epoch: 4,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
			},
			validatorIDs: validatorIDs,
		},
		{
			name:  "one step epoch exact",
			epoch: 5,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
			},
			validatorIDs: validatorIDs[:2],
		},
		{
			name:  "one step epoch after",
			epoch: 6,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
			},
			validatorIDs: validatorIDs[:2],
		},
		{
			name:  "two steps epoch before first",
			epoch: 4,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
				{Epoch: 10, Cutoff: 1},
			},
			validatorIDs: validatorIDs,
		},
		{
			name:  "two steps epoch exact first",
			epoch: 5,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
				{Epoch: 10, Cutoff: 1},
			},
			validatorIDs: validatorIDs[:2],
		},
		{
			name:  "two steps epoch after first",
			epoch: 6,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
				{Epoch: 10, Cutoff: 1},
			},
			validatorIDs: validatorIDs[:2],
		},
		{
			name:  "two steps epoch before second",
			epoch: 9,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
				{Epoch: 10, Cutoff: 1},
			},
			validatorIDs: validatorIDs[:2],
		},
		{
			name:  "two steps epoch exact second",
			epoch: 10,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
				{Epoch: 10, Cutoff: 1},
			},
			validatorIDs: validatorIDs[:1],
		},
		{
			name:  "two steps epoch after second",
			epoch: 11,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
				{Epoch: 10, Cutoff: 1},
			},
			validatorIDs: validatorIDs[:1],
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			v := ValidatorsDefault{
				validators: validators,
				steps:      test.steps,
			}

			gotValidators, err := v.ByEpoch(test.epoch)
			require.NoError(t, err)

			assert.Len(t, gotValidators, len(test.validatorIDs))
			for _, validatorID := range test.validatorIDs {
				assert.Contains(t, gotValidators, validatorID)
			}
		})
	}
}
