package config

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

var (
	configRoot    = "/etc/minstor"
	obsConfigName = "obs.toml"
)

type ObsConfig struct {
	Devices   string `toml:"devices"`
	LogLevel  string `toml:"log_level"`
	LogCaller bool   `toml:"log_caller"`
}

func SetConfigRoot(dir string) {
	configRoot = dir
}

func SetObsConfigName(name string) {
	obsConfigName = name
}

func NewObsConfig() (*ObsConfig, error) {
	content, err := os.ReadFile(filepath.Join(configRoot, obsConfigName))
	if err != nil {
		return nil, err
	}

	conf := ObsConfig{
		Devices:   "/srv/minstor",
		LogLevel:  "Info",
		LogCaller: false,
	}

	err = toml.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
