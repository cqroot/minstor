package config

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

var (
	configRoot             = "/etc/minstor"
	objectServerConfigName = "object-server.toml"
)

type ObjectServerConfig struct {
	Devices  string `toml:"devices"`
	LogLevel string `toml:"log_level"`
}

func SetConfigRoot(dir string) {
	configRoot = dir
}

func SetObjectServerConfigName(name string) {
	objectServerConfigName = name
}

func NewObjectServerConfig() (*ObjectServerConfig, error) {
	content, err := os.ReadFile(filepath.Join(configRoot, objectServerConfigName))
	if err != nil {
		return nil, err
	}

	conf := ObjectServerConfig{
		Devices:  "/srv/minstor",
		LogLevel: "Info",
	}

	err = toml.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
