package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var API_KEY string = ""

var CLIENT *http.Client = &http.Client{}

func getWeatherForecast(lat float64, long float64) WeatherData {
	api := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v&units=metric",lat, long, API_KEY)

	resp, err := CLIENT.Get(api)
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

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	API_KEY = os.Getenv("API_KEY")

	city := getUserCity()


	locations := getGeoCodes(city)
	
	fmt.Println(locations)
	filtered_locations := filterLocations(city, locations)

	choice := getLocationIndex(filtered_locations)

	lat, lon := getCoordinates(filtered_locations, choice)

	x := getWeatherForecast(lat, lon)

	printForecast(x)
}
