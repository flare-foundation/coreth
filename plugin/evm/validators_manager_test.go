// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/database"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"
)

type retrieverMock struct {
	ByEpochFunc func(epoch uint64) (map[ids.ShortID]uint64, error)
}

func (r retrieverMock) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {
	return r.ByEpochFunc(epoch)
}

func TestNewValidatorsManager(t *testing.T) {

	defaultRetriever := &retrieverMock{}
	ftsoRetriever := &retrieverMock{}
	activeRetriever := &retrieverMock{}
	transitionRetriever := &retrieverMock{}

	got := NewValidatorsManager(logging.NoLog{}, defaultRetriever, ftsoRetriever, activeRetriever, transitionRetriever)

	require.NotNil(t, got)
	assert.Equal(t, defaultRetriever, got.defaultValidators)
	assert.Equal(t, ftsoRetriever, got.ftsoValidators)
	assert.Equal(t, activeRetriever, got.activeValidators)
	assert.Equal(t, transitionRetriever, got.transitionValidators)
}

func TestValidatorsManager_ActiveValidators(t *testing.T) {

	testEpoch := uint64(1)
	wantValidators := map[ids.ShortID]uint64{
		{13}: 37,
	}
	t.Run("nominal case", func(t *testing.T) {
		t.Parallel()

		testRetriever := &retrieverMock{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				assert.Equal(t, testEpoch, epoch)
				return wantValidators, nil
			},
		}
		testTransitioner := &retrieverMock{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				return nil, errors.New("dummy error")
			},
		}

		subject := &ValidatorsManager{
			log:                  logging.NoLog{},
			activeValidators:     testRetriever,
			transitionValidators: testTransitioner,
		}

		gotValidators, err := subject.ActiveValidators(testEpoch)
		require.NoError(t, err)
		assert.Equal(t, wantValidators, gotValidators)
	})

	t.Run("gracefully falls back on transitioner", func(t *testing.T) {
		t.Parallel()

		testRetriever := &retrieverMock{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				assert.Equal(t, testEpoch, epoch)
				return nil, database.ErrNotFound
			},
		}
		testTransitioner := &retrieverMock{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				return wantValidators, nil
			},
		}

		subject := &ValidatorsManager{
			log:                  logging.NoLog{},
			activeValidators:     testRetriever,
			transitionValidators: testTransitioner,
		}

		gotValidators, err := subject.ActiveValidators(testEpoch)
		require.NoError(t, err)
		assert.Equal(t, wantValidators, gotValidators)
	})

	t.Run("handles failure to retrieve active validators", func(t *testing.T) {
		t.Parallel()

		testRetriever := &retrieverMock{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				assert.Equal(t, testEpoch, epoch)
				return nil, errors.New("dummy error")
			},
		}
		testTransitioner := &retrieverMock{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				return wantValidators, nil
			},
		}

		subject := &ValidatorsManager{
			log:                  logging.NoLog{},
			activeValidators:     testRetriever,
			transitionValidators: testTransitioner,
		}

		_, err := subject.ActiveValidators(testEpoch)
		assert.Error(t, err)
	})

	t.Run("handles failure to retrieve active validators", func(t *testing.T) {
		t.Parallel()

		testRetriever := &retrieverMock{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				assert.Equal(t, testEpoch, epoch)
				return nil, database.ErrNotFound
			},
		}
		testTransitioner := &retrieverMock{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				return nil, errors.New("dummy error")
			},
		}

		subject := &ValidatorsManager{
			log:                  logging.NoLog{},
			activeValidators:     testRetriever,
			transitionValidators: testTransitioner,
		}

		_, err := subject.ActiveValidators(testEpoch)
		assert.Error(t, err)
	})
}
