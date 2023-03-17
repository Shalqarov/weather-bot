package weather

import (
	"fmt"
	"math"

	owm "github.com/briandowns/openweathermap"
)

func CurrentTemperature(location string, apiKey string) (string, error) {
	w, err := owm.NewCurrent("C", "ru", apiKey)
	if err != nil {
		return "", err
	}

	err = w.CurrentByName("Astana")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Локация: %s\nТемпература воздуха: %.f°C",
		location, math.Round(w.Main.Temp),
	), nil
}
