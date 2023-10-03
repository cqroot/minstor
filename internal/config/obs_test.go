package config_test

import (
	"testing"

	"github.com/cqroot/minstor/internal/config"
	"github.com/stretchr/testify/require"
)

func TestDefaultObsConfig(t *testing.T) {
	config.SetConfigRoot("./testdata")
	config.SetObsConfigName("empty-obs.toml")

	conf, err := config.NewObsConfig()
	require.Nil(t, err)

	require.Equal(t, config.ObsConfig{
		Devices:   "/srv/minstor",
		LogLevel:  "Info",
		LogCaller: false,
	}, *conf)
}

func TestObsConfig(t *testing.T) {
	config.SetConfigRoot("./testdata")
	config.SetObsConfigName("obs.toml")

	conf, err := config.NewObsConfig()
	require.Nil(t, err)

	require.Equal(t, config.ObsConfig{
		Devices:   "/data/minstor",
		LogLevel:  "Debug",
		LogCaller: true,
	}, *conf)
}
