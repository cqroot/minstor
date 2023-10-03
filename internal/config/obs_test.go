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
		FieldBindIp:    "127.0.0.1",
		FieldBindPort:  10001,
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
		FieldBindIp:    "8.8.8.8",
		FieldBindPort:  10002,
	}, *conf)
}
