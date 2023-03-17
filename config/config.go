package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

const configPath string = "./config.json"

type Config struct {
	WeatherAPI  string `json:"weather_api"`
	Port        string `json:"port"`
	TelegramAPI string `json:"telegram_api"`
}

var (
	cfg   *Config
	cOnce sync.Once
)

func GetConfig() *Config {
	cOnce.Do(func() {
		data, err := os.ReadFile(configPath)
		if err != nil {
			log.Fatalln(err.Error())
		}
		if err = json.Unmarshal(data, &cfg); err != nil {
			log.Fatalln(err.Error())
		}
	})
	return cfg
}
