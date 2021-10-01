package flare

import (
	"errors"
)

var (
	ErrRejected = errors.New("rejected")
	ErrUnknown  = errors.New("unknown")
)
