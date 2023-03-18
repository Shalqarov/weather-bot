package main

import (
	"github.com/Shalqarov/weather-bot/config"
	initialize "github.com/Shalqarov/weather-bot/init"
	"github.com/Shalqarov/weather-bot/internal/app"
	"github.com/Shalqarov/weather-bot/internal/bot"
	"github.com/Shalqarov/weather-bot/internal/database/postgres"
	"github.com/Shalqarov/weather-bot/pkg"
)

const configPath = "./config.toml"

func init() {
	// Config init
	config.GetConfig()
	initialize.Migration("./internal/database/postgres/migrations/")
	// db init
	postgres.GetDB()
	// users init
	pkg.UsersMap()
	// bot init
	bot.GetBot()
}

func main() {
	app.Run()
}
