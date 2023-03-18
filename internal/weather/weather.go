package weather

import (
	"fmt"
	"math"

	owm "github.com/briandowns/openweathermap"

	"github.com/Shalqarov/weather-bot/config"
)

func CurrentTemperature(location string) (string, error) {
	w, err := owm.NewCurrent("C", "ru", config.GetConfig().WeatherAPI)
	if err != nil {
		return "", err
	}
	err = w.CurrentByName(location)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"Локация: %s\n"+
			"Температура воздуха: %.f°C, ощущается как %.f°C\n"+
			"Влажность воздуха: %d\n"+
			"Небо: %s\n",
		location,
		math.Round(w.Main.Temp),
		math.Round(w.Main.FeelsLike),
		w.Main.Humidity,
		w.Weather[0].Description,
	), nil
}
