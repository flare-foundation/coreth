package valstate

import (
	"errors"
)

var (
	ErrNoEntries = errors.New("no entries")
	ErrNotFound  = errors.New("not found")
)
