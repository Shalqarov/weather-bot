package weather

import (
	"fmt"
	"math"

	"github.com/Shalqarov/weather-bot/config"
	owm "github.com/briandowns/openweathermap"
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
