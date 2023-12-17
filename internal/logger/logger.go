package logger

import (
	"os"

	"golang.org/x/exp/slog"
)

var logger *slog.Logger

func Init() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug.Level(),
	}))
}

func Debug(msg string, args ...any) {
	logger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}

func ErrorE(msg string, err error, args ...any) {
	args = append(args, slog.String("error", err.Error()))
	logger.Error(msg, args...)
}
