package weather_client

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func GetWeather(latitude float64, longitude float64) (string, error) {
	client := resty.New()
	url := fmt.Sprintf("https://api.weather.gov/points/%f,%f", latitude, longitude)
	resp, err := client.R().Get(url)
	if err != nil {
		return "", fmt.Errorf("error getting weather response URL: %w", err)
	}

	fmt.Printf("First API response status: %d\n", resp.StatusCode())

	var weatherResponse WeatherResponse
	err = json.Unmarshal(resp.Body(), &weatherResponse)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling weather response: %w", err)
	}

	forecastURL := weatherResponse.Properties.Forecast
	fmt.Printf("Forecast URL: %s\n", forecastURL)

	if forecastURL == "" {
		return "", fmt.Errorf("forecast URL is empty")
	}

	forecastResp, err := client.R().Get(forecastURL)
	if err != nil {
		return "", fmt.Errorf("error getting forecast response URL: %w", err)
	}

	return string(forecastResp.Body()), nil
}
