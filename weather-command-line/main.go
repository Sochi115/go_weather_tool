package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "github.com/gin-gonic/gin"
)

var API_KEY string = "c7879271fb0d4d2004f8828467eeeb4f"
var CLIENT *http.Client = &http.Client{}

// func helloWorld(c *gin.Context) {
// 	fmt.Fprintf(c.Writer, "Hello World")
// }

func getGeoCode(country string) (float64, float64) {
	api := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%v&limit=5&apikey=%v", country, API_KEY)

	resp, err := CLIENT.Get(api)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var x interface{}

	json.Unmarshal(body, &x)

	locations := x.([]interface{})

	area := locations[0]
	areaMap := area.(map[string]interface{})

	lat := areaMap["lat"]
	lon := areaMap["lon"]
	return lat.(float64), lon.(float64)
}

func getWeatherForecast(lat float64, long float64) WeatherData {
	api := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v",lat, long, API_KEY)

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
	// lat, lon := getGeoCode("Phnom Penh")
	lat := 11.5682711
	lon :=  104.9224426
	x := getWeatherForecast(lat, lon)

	fmt.Println(x)
}
