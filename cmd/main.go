package main

import (
	initialize "github.com/Shalqarov/weather-bot/init"
	"github.com/Shalqarov/weather-bot/internal/app"
)

func init() {
	initialize.Config()
	initialize.Migration()
	initialize.DB()
	initialize.Bot()
}

func main() {
	app.Run()
}
