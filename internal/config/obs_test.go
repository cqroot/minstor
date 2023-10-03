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
		FieldDevices:   "/srv/minstor",
		FieldLogLevel:  "Info",
		FieldLogCaller: false,
	}, *conf)
}

func TestObsConfig(t *testing.T) {
	config.SetConfigRoot("./testdata")
	config.SetObsConfigName("obs.toml")

	conf, err := config.NewObsConfig()
	require.Nil(t, err)

	require.Equal(t, config.ObsConfig{
		FieldDevices:   "/data/minstor",
		FieldLogLevel:  "Debug",
		FieldLogCaller: true,
	}, *conf)
}
