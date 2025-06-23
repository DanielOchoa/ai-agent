package tools

import (
	weather_client "agent/client/weather"
	"encoding/json"
	"fmt"
)

var WeatherDefinition = ToolDefinition{
	Name:        "weather",
	Description: "Get the weather for a given latitude and longitude. Do not tell user you are using latitude/longitude. Use a funny demeanor when replying with the weather. Includes forecast so you can tell how it will be tomorrow for example. This will only work for places in the US. If asked for a place outside the US, say you don't know the weather there.",
	InputSchema: WeatherInputSchema,
	Function:    Weather,
}

type WeatherInput struct {
	Latitude  float64 `json:"latitude" jsonschema_description:"The latitude of the location to get the weather for."`
	Longitude float64 `json:"longitude" jsonschema_description:"The longitude of the location to get the weather for."`
}

var WeatherInputSchema = GenerateSchema[WeatherInput]()

func Weather(input json.RawMessage) (string, error) {
	weatherInput := WeatherInput{}
	err := json.Unmarshal(input, &weatherInput)
	if err != nil {
		panic(err)
	}

	weatherResp, err := weather_client.GetWeather(weatherInput.Latitude, weatherInput.Longitude)
	if err != nil {
		return "", fmt.Errorf("error getting weather from weather client: %w", err)
	}

	return weatherResp, nil
}
