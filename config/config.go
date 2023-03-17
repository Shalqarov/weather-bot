package config

import (
	"github.com/BurntSushi/toml"
)

var (
	cfg *Config
)

type Config struct {
	WeatherAPI  string
	Port        string
	TelegramAPI string
}

func GetConfig() *Config {
	return cfg
}

func NewConfig(path string) error {
	cfg := &Config{}
	_, err := toml.DecodeFile(path, cfg)
	return err
}
