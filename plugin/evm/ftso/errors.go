package ftso

import (
	"errors"
)

var (
	ErrNoPriceSubmitter = errors.New("no genesis price submitter")
	ErrFTSONotDeployed  = errors.New("FTSO system not deployed")
	ErrFTSONotActive    = errors.New("FTSO system not activated")
)
