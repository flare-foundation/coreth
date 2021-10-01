package flare

import (
	"fmt"
)

func Hash(data []byte) [32]byte {
	if len(data) < 32 {
		panic(fmt.Sprintf("insufficient data for hash (%d < %d)", len(data), 32))
	}
	var hash [32]byte
	copy(hash[:], data)
	return hash
}
