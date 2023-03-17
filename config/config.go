package config

import "github.com/BurntSushi/toml"

type Config struct {
	APIKey string
	Port   string
}

func NewConfig(path string) (*Config, error) {
	c := &Config{}
	_, err := toml.DecodeFile(path, c)
	return c, err
}
