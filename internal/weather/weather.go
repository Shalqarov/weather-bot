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
		"Локация: %s\nТемпература воздуха: %.f°C",
		location, math.Round(w.Main.Temp),
	), nil
}
