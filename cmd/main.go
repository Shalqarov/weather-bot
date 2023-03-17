package main

import (
	"log"

	"github.com/Shalqarov/weather-bot/config"
	"github.com/Shalqarov/weather-bot/internal/app"
)

const configPath = "./config.toml"

func main() {
	c, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	app.Run(c)
}
