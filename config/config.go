package config

import "github.com/BurntSushi/toml"

type Config struct {
	WeatherAPI  string
	Port        string
	TelegramAPI string
}

func NewConfig(path string) (*Config, error) {
	c := &Config{}
	_, err := toml.DecodeFile(path, c)
	return c, err
}
