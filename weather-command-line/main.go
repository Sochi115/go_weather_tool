package main

import (
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

func getGeoCode(country string) string{
	api := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%v&limit=5&apikey=%v", country, API_KEY )

	resp , err := CLIENT.Get(api)
	if err!= nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err!= nil {
		fmt.Println(err.Error())
	}

	return string(body)

}
// func getWeatherForecast(lat float64, long float64){
// 	api := fmt.Sprintf("api.openweathermap.org/data/2.5/forecast/daily?lat=%v&lon=%v&cnt=3&appid=%v",lat, long, API_KEY)

// }

func main() {
	resp := getGeoCode("Phnom Penh")
	fmt.Println(resp)
}