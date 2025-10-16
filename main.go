package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type apiConfigData struct {
	OpenWeatherMapKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData

	err = json.Unmarshal(data, &c)
	if err != nil {
		return apiConfigData{}, err
	}

	return c, nil
}
