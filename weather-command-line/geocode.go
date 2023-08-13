package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func getGeoCodes(country string) []interface{} {
	country = strings.TrimSpace(country)
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

	return locations
}
func filterLocations(city string, locations[]interface{}) []interface{} {
	var result []interface{}
	city = strings.TrimRight(city,"\r\n")


	for _, v := range(locations) {
			area := v.(map[string]interface{})

			x := area["name"]
			name := x.(string)

			if name == city {
				result = append(result, v)
			}
	}
	return result
}

func getCoordinates(locations []interface{}, index int16) (float64,float64){

	area := locations[0]
	areaMap := area.(map[string]interface{})

	lat := areaMap["lat"]
	lon := areaMap["lon"]
	return lat.(float64), lon.(float64)
}