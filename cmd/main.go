package main

import (
	"log"

	"github.com/Shalqarov/weather-bot/config"
	"github.com/Shalqarov/weather-bot/internal/app"
)

func main() {
	c, err := config.NewConfig("./config.toml")
	if err != nil {
		log.Fatalln(err)
	}
	app.Run(c)
}
