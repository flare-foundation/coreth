package flare

import (
	"github.com/rs/zerolog"
)

func Log() zerolog.Logger {
	return system.log
}
