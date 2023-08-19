package main

import (
	"log"
	"weather-command-line/weather-command-line/userprompt"
)

func main() {

	httpClient := getHttpClient()

	choice := promptZipOrCityChoice()
	if choice == 2 {
		promptByZip(httpClient)
	} else {
		promptByCity(httpClient)
	}
}

func promptByZip(httpClient *client) {
	var zipprompt = userprompt.Zipprompt{}
	var _ userprompt.Prompt = userprompt.Zipprompt{}

	lat, long, err := zipprompt.GetCoordinates(httpClient.client, httpClient.apiKey)

	if err != nil {
		log.Default().Println(err.Error())
		promptByCity(httpClient)
		return
	}
	displayWeatherData(lat, long)
}

func promptByCity(httpClient *client) {
	var cityprompt = userprompt.Cityprompt{}
	lat, long, err := cityprompt.GetCoordinates(httpClient.client, clientInstance.apiKey)

	if err != nil {
		log.Fatal(err.Error())
	}

	displayWeatherData(lat, long)
}
