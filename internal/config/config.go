package config

type Config interface {
	LogLevel() string
	LogCaller() bool
}
