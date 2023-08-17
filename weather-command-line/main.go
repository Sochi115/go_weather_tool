package main

import (
	"weather-command-line/weather-command-line/userprompt"
)

func main() {

	httpClient := getHttpClient()
	var cityprompt = userprompt.Cityprompt{}
	lat, long := cityprompt.GetCoordinates(httpClient.client, clientInstance.apiKey)
	displayWeatherData(lat, long)
}
