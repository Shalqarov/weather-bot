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
	UserDB      string `json:"user_db"`
	PasswordDB  string `json:"password_db"`
	HostDB      string `json:"host_db"`
	PortDB      string `json:"port_db"`
	NameDB      string `json:"name_db"`
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
