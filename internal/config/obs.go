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
	FieldLogLevel  string `toml:"log_level"`
	FieldLogCaller bool   `toml:"log_caller"`
	FieldDevices   string `toml:"devices"`
	FieldBindIp    string `toml:"bind_ip"`
	FieldBindPort  int    `toml:"bind_port"`
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
		FieldDevices:   "/srv/minstor",
		FieldLogLevel:  "Info",
		FieldLogCaller: false,
		FieldBindIp:    "127.0.0.1",
		FieldBindPort:  10001,
	}

	err = toml.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func (c ObsConfig) LogLevel() string {
	return c.FieldLogLevel
}

func (c ObsConfig) LogCaller() bool {
	return c.FieldLogCaller
}

func (c ObsConfig) Devices() string {
	return c.FieldDevices
}

func (c ObsConfig) BindIp() string {
	return c.FieldBindIp
}

func (c ObsConfig) BindPort() int {
	return c.FieldBindPort
}
