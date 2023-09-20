package config_test

import (
	"testing"

	"github.com/cqroot/minstor/internal/config"
	"github.com/stretchr/testify/require"
)

func TestDefaultObjectServerConfig(t *testing.T) {
	config.SetConfigRoot("./testdata")
	config.SetObjectServerConfigName("empty-object-server.toml")

	conf, err := config.NewObjectServerConfig()
	require.Nil(t, err)

	require.Equal(t, config.ObjectServerConfig{
		Devices:   "/srv/minstor",
		LogLevel:  "Info",
		LogCaller: false,
	}, *conf)
}

func TestObjectServerConfig(t *testing.T) {
	config.SetConfigRoot("./testdata")
	config.SetObjectServerConfigName("object-server.toml")

	conf, err := config.NewObjectServerConfig()
	require.Nil(t, err)

	require.Equal(t, config.ObjectServerConfig{
		Devices:   "/data/minstor",
		LogLevel:  "Debug",
		LogCaller: true,
	}, *conf)
}
