package weather_client

// https://api.weather.gov/points/{latitude},{longitude} <- URL to get location data for lat/long.
// { properties.relativeLocation.properties.city, state } <- City and state for the location.
// { properties.forecast } <- URL to get the forecast for the location.
type WeatherResponse struct {
	Properties struct {
		RelativeLocation struct {
			Properties struct {
				City  string `json:"city"`
				State string `json:"state"`
			} `json:"properties"`
		} `json:"relativeLocation"`
		Forecast string `json:"forecast"`
	} `json:"properties"`
}

// from forecast URL:
// { properties.periods } <- Array of periods.
// { properties.periods.number } <- The number of the period.
// { properties.periods.name } <- The name of the period.
type WeatherForecastResponse struct {
	Properties struct {
		Periods []struct {
			Number int    `json:"number"`
			Name   string `json:"name"`
		} `json:"periods"`
	} `json:"properties"`
}
