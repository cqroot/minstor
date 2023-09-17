package config

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

var configRoot = "/etc/minstor"

type ObjectServerConfig struct {
	Devices string
}

func SetConfigRoot(dir string) {
	configRoot = dir
}

func ReadObjectServerConfig() (*ObjectServerConfig, error) {
	content, err := os.ReadFile(filepath.Join(configRoot, "/object-server.toml"))
	if err != nil {
		return nil, err
	}

	config := ObjectServerConfig{}

	err = toml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
