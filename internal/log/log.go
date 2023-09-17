package log

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct{}

func New() zerolog.Logger {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.999"
	logger := zerolog.New(os.Stderr).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	return logger
}
