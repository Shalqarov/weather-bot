package app

import (
	"fmt"
	"log"

	"github.com/Shalqarov/weather-bot/config"
	"github.com/Shalqarov/weather-bot/internal/weather"
)

func Run(config *config.Config) {
	w, err := weather.CurrentTemperature("Astana", config.WeatherAPI)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(w)
}
