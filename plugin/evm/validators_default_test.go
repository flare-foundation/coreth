package evm

import (
	"math/big"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/params"
)

func TestNewValidatorsDefault(t *testing.T) {
	tests := []struct {
		name    string
		chainID *big.Int
		weight  uint64
		amount  int
		wantErr require.ErrorAssertionFunc
	}{
		{
			chainID: params.CostonChainID,
			weight:  costonValidatorWeight,
			amount:  5,
			wantErr: require.NoError,
		},
		{
			chainID: params.SongbirdChainID,
			weight:  songbirdValidatorWeight,
			amount:  20,
			wantErr: require.NoError,
		},
		{
			chainID: params.FlareChainID,
			weight:  flareValidatorWeight,
			amount:  20,
			wantErr: require.Error,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewValidatorsDefault(test.chainID)
			test.wantErr(t, err)

			if err == nil {
				assert.Len(t, got.validators, test.amount)
				gotWeight := uint64(0)
				for _, v := range got.validators {
					gotWeight += v
				}
				assert.Equal(t, uint64(test.amount)*test.weight, gotWeight)
			}
		})
	}
}

func TestGetDefaultValidators_CustomValidators(t *testing.T) {
	err := os.Setenv("CUSTOM_VALIDATORS", "NodeID-3M9KVT6ixi4gVMisbm5TnPXYXgFN5LHuv,NodeID-NnX4fajAmyvpL9RLfheNdc47FKKDuQW8i")
	require.NoError(t, err)

	got, err := NewValidatorsDefault(big.NewInt(444))
	require.NoError(t, err)

	assert.Len(t, got.validators, 2)
	var gotWeight uint64
	for _, v := range got.validators {
		gotWeight += v
	}
	assert.Equal(t, uint64(2)*customValidatorWeight, gotWeight)
}

func TestValidatorsDefault_ByEpoch(t *testing.T) {

	t.Run("nominal case", func(t *testing.T) {
		t.Parallel()

		var totalWeight uint64
		var epoch uint64 = 12

		v, err := NewValidatorsDefault(params.CostonChainID)
		require.NoError(t, err)

		got, err := v.ByEpoch(epoch)
		require.NoError(t, err)

		for _, u := range got {
			totalWeight += u
		}

		assert.Equal(t, uint64(5*costonValidatorWeight), totalWeight)

		assert.Len(t, got, 5)

	})

	t.Run("edge cases", func(t *testing.T) {
		t.Parallel()

		var totalWeight uint64
		var epoch uint64 = 1604

		v, err := NewValidatorsDefault(params.CostonChainID)
		require.NoError(t, err)

		got, err := v.ByEpoch(epoch)
		require.NoError(t, err)

		for _, u := range got {
			totalWeight += u
		}

		assert.Equal(t, uint64(4*costonValidatorWeight), totalWeight)

		assert.Len(t, got, 4)

		epoch = 1772
		got, err = v.ByEpoch(epoch)
		require.NoError(t, err)
		assert.Len(t, got, 3)

	})
}
