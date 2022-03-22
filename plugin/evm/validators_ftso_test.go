package evm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithRootDegree(t *testing.T) {
	wantDegree := uint(35)

	ftsoCfg := FTSOConfig{}
	WithRootDegree(wantDegree)(&ftsoCfg)

	assert.Equal(t, wantDegree, ftsoCfg.RootDegree)
}
