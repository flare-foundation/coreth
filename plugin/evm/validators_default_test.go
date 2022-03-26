package evm

import (
	"math/big"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/params"
)

func TestGetDefaultValidators(t *testing.T) {
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

			got, err := getDefaultValidators(test.chainID)
			test.wantErr(t, err)

			if err == nil {
				assert.Len(t, got, test.amount)
				gotWeight := uint64(0)
				for _, v := range got {
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

	got, err := getDefaultValidators(big.NewInt(444))
	require.NoError(t, err)

	assert.Len(t, got, 2)
	var gotWeight uint64
	for _, v := range got {
		gotWeight += v
	}
	assert.Equal(t, uint64(2)*customValidatorWeight, gotWeight)
}
