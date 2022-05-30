package validatordb

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/ids"

	"github.com/flare-foundation/coreth/core/rawdb"
	"github.com/flare-foundation/coreth/core/state"
	"github.com/flare-foundation/coreth/params"
)

func TestState_Epoch(t *testing.T) {
	st, db := testState(t)

	epoch := uint64(10)

	err := st.SetEpoch(epoch)
	require.NoError(t, err)

	res, err := st.GetEpoch()
	require.NoError(t, err)
	assert.Equal(t, epoch, res)

	// commit the changes
	root, err := st.RootHash()
	require.NoError(t, err)
	st = commitAndRecreate(t, root, db)

	res, err = st.GetEpoch()
	require.NoError(t, err)
	assert.Equal(t, epoch, res)
}

func TestState_Mapping(t *testing.T) {
	st, db := testState(t)

	providerAddr1 := common.BytesToAddress([]byte("bandera"))
	providerID1 := ids.GenerateTestShortID()
	providerAddr2 := common.BytesToAddress([]byte("stetsko"))
	providerID2 := ids.GenerateTestShortID()

	err := st.SetMapping(providerAddr1, providerID1)
	require.NoError(t, err)
	err = st.SetMapping(providerAddr2, providerID2)
	require.NoError(t, err)

	resID, err := st.GetMapping(providerAddr1)
	require.NoError(t, err)
	assert.Equal(t, providerID1, resID)

	resID, err = st.GetMapping(providerAddr2)
	require.NoError(t, err)
	assert.Equal(t, providerID2, resID)

	resID, err = st.GetMapping(common.BytesToAddress([]byte("unknown")))
	require.Error(t, err)
	require.Equal(t, ids.ShortID{}, resID)

	resMapping, err := st.AllMappings()
	require.NoError(t, err)
	assert.Len(t, resMapping, 2)
	assert.Equal(t, providerID1, resMapping[providerAddr1])
	assert.Equal(t, providerID2, resMapping[providerAddr2])

	// commit the changes
	root, err := st.RootHash()
	require.NoError(t, err)
	st = commitAndRecreate(t, root, db)

	resMapping, err = st.AllMappings()
	require.NoError(t, err)
	assert.Len(t, resMapping, 2)
	assert.Equal(t, providerID1, resMapping[providerAddr1])
	assert.Equal(t, providerID2, resMapping[providerAddr2])
}

func TestState_Candidates(t *testing.T) {
	st, db := testState(t)

	candidates := []*Candidate{
		{
			Providers: []common.Address{common.BytesToAddress([]byte("sirko")), common.BytesToAddress([]byte("dashkevych"))},
			NodeID:    ids.GenerateTestShortID(),
			Votepower: 100,
		},
		{
			Providers: []common.Address{common.BytesToAddress([]byte("apostol")), common.BytesToAddress([]byte("sulyma"))},
			NodeID:    ids.GenerateTestShortID(),
			Votepower: 200,
		},
	}

	err := st.SetCandidates(candidates)
	require.NoError(t, err)

	res, err := st.GetCandidates()
	require.NoError(t, err)
	assert.Equal(t, candidates, res)

	// commit the changes
	root, err := st.RootHash()
	require.NoError(t, err)
	st = commitAndRecreate(t, root, db)

	res, err = st.GetCandidates()
	require.NoError(t, err)
	assert.Equal(t, candidates, res)
}

func TestState_Validators(t *testing.T) {
	st, db := testState(t)

	validators := []*Validator{
		{
			Providers: []common.Address{common.BytesToAddress([]byte("lysenko")), common.BytesToAddress([]byte("leontovych"))},
			NodeID:    ids.GenerateTestShortID(),
			Weight:    20,
		},
		{
			Providers: []common.Address{common.BytesToAddress([]byte("skoryk")), common.BytesToAddress([]byte("ivasyuk"))},
			NodeID:    ids.GenerateTestShortID(),
			Weight:    30,
		},
	}

	err := st.SetValidators(validators)
	require.NoError(t, err)

	res, err := st.GetValidators()
	require.NoError(t, err)
	assert.Equal(t, validators, res)

	// commit the changes
	root, err := st.RootHash()
	require.NoError(t, err)
	st = commitAndRecreate(t, root, db)

	res, err = st.GetValidators()
	require.NoError(t, err)
	assert.Equal(t, validators, res)
}

func commitAndRecreate(t *testing.T, root common.Hash, st *state.StateDB) *State {
	t.Helper()

	// store the changes into the repo
	st.SetCode(params.ValidationAddress, root[:])

	_, err := st.Commit(true)
	require.NoError(t, err)

	repo := NewRepository(nil)
	repo.db = st.Database()

	// get current repo
	code := st.GetCode(params.ValidationAddress)
	s, err := repo.WithRoot(common.BytesToHash(code))
	require.NoError(t, err)

	return s
}

func testState(t *testing.T) (s *State, db *state.StateDB) {
	t.Helper()

	mdb := rawdb.NewMemoryDatabase()
	db, err := state.New(common.Hash{}, state.NewDatabase(mdb), nil)
	require.NoError(t, err)

	hash, err := db.Commit(true)
	require.NoError(t, err)

	repo := NewRepository(nil)
	repo.db = db.Database()

	s, err = repo.WithRoot(hash)
	require.NoError(t, err)

	return
}
