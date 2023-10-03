package logger

import (
	"os"
	"strings"

	"github.com/cqroot/minstor/internal/config"
	"github.com/rs/zerolog"
)

type Logger = zerolog.Logger

func logLevelFromString(logLevel string) zerolog.Level {
	logLevel = strings.ToLower(logLevel)
	switch logLevel {
	case "error":
		return zerolog.WarnLevel
	case "warn":
		return zerolog.WarnLevel
	case "debug":
		return zerolog.DebugLevel
	}
	return zerolog.InfoLevel
}

func New(conf *config.ObsConfig) Logger {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.999"
	l := zerolog.New(os.Stderr).
		Level(logLevelFromString(conf.LogLevel)).
		With().
		Timestamp()

	if conf.LogCaller {
		return l.Caller().Logger()
	}
	return l.Logger()
}
