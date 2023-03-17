package main

import (
	"github.com/Shalqarov/weather-bot/config"
	"github.com/Shalqarov/weather-bot/internal/app"
)

const configPath = "./config.toml"

func init() {
	// Config initialization
	config.GetConfig()
}

func main() {
	app.Run()
}
