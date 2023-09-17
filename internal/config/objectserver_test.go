package config_test

import (
	"testing"

	"github.com/cqroot/minstor/internal/config"
	"github.com/stretchr/testify/require"
)

func TestObjectServerConfig(t *testing.T) {
	config.SetConfigRoot("./testdata")

	conf, err := config.ReadObjectServerConfig()
	require.Nil(t, err)

	require.Equal(t, config.ObjectServerConfig{
		Devices: "/srv/minstor",
	}, *conf)
}
