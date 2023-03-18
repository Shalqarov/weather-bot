package main

import (
	"github.com/Shalqarov/weather-bot/config"
	"github.com/Shalqarov/weather-bot/internal/app"
	"github.com/Shalqarov/weather-bot/internal/bot"
	"github.com/Shalqarov/weather-bot/pkg"
)

const configPath = "./config.toml"

func init() {
	// Config init
	config.GetConfig()
	// users init
	pkg.UsersMap()
	// bot init
	bot.GetBot()
}

func main() {
	app.Run()
}
