package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func getCityCoordinates(city string, apiKey string) (float64, float64) {
	var locations = getGeoCodesData(city, apiKey)
	filtered_locations := filterLocations(city, locations)

	index := promptLocationChoice(filtered_locations)

	chosenLocation := filtered_locations[index]
	return extractCoordinates(chosenLocation.(map[string]interface{}))
}

func getGeoCodesData(city string, apiKey string) []interface{} {
	city = strings.TrimSpace(city)
	api := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%v&limit=5&apikey=%v", city, apiKey )

	resp, err := CLIENT.Get(api)
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
func filterLocations(city string, locations []interface{}) []interface{} {
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

func extractCoordinates(location map[string]interface{}) (float64, float64) {
	lat := location["lat"]
	lon := location["lon"]
	return lat.(float64), lon.(float64)
}
