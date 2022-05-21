package server

import (
	"encoding/json"
	"net/http"
)

type Weather struct {
	City     string `json:"city"`
	Forecast string `json:"forecast"`
}

func GetWeather(url string) (*Weather, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var weather Weather
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}
