package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func displayWeatherData(lat float64, long float64) {
	weatherData := queryWeatherForecast(lat, long)
	printForecast(weatherData)
}
func queryWeatherForecast(lat float64, long float64) WeatherData{
	httpClient := getHttpClient()

	api := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v&units=metric", lat, long, httpClient.apiKey)

	resp, err := httpClient.client.Get(api)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	x := WeatherData{}
	json.Unmarshal(body, &x)
	return x
}
