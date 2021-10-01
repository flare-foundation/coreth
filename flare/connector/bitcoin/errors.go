package bitcoin

import (
	"errors"
)

var (
	ErrOutputNotFound    = errors.New("output not found")
	ErrInvalidKeyType    = errors.New("invalid key type")
	ErrTooManyRecipients = errors.New("too many recipients")
)
