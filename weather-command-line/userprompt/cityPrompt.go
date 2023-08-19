package userprompt

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Cityprompt struct {
}

func (c Cityprompt) GetCoordinates(httpClient *http.Client, apiKey string) (float64, float64, error) {
	city := promptCity()
	countryCode := promptCountryCode()

	api := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%v,,%v&limit=5&apikey=%v", city, countryCode, apiKey)

	locations := getCityList(httpClient, api)

	if len(locations) > 1 {
		locations = filterCityList(city, locations)
	}

	index := promptLocationChoice(locations)
	chosenLocation := locations[index]

	lat, lon :=  extractCoordinates(chosenLocation.(map[string]interface{}))

	return lat, lon, nil
}
func promptCity() string {
	fmt.Print("Enter city: ")
	city, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	city = strings.TrimSpace(city)
	return strings.TrimRight(city, "\r\n")
}

func getCityList(httpClient *http.Client, apiUrl string) []interface{} {
	resp, err := httpClient.Get(apiUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var x interface{}

	json.Unmarshal(body, &x)

	locations := x.([]interface{})

	return locations
}

func filterCityList(city string, locations []interface{}) []interface{} {
	var result []interface{}
	city = strings.TrimRight(city, "\r\n")

	for _, v := range locations {
		area := v.(map[string]interface{})

		x := area["name"]
		name := x.(string)

		if name == city {
			result = append(result, v)
		}
	}
	return result
}
