package userprompt

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Cityprompt struct {
}

var reader = getReaderInstance()

func (c Cityprompt) GetCoordinates(httpClient *http.Client, apiKey string) (float64, float64) {
	city := promptCity()
	countryCode := promptCountryCode()

	api := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%v,,%v&limit=5&apikey=%v", strings.TrimSpace(city), strings.TrimSpace(countryCode), apiKey)

	locations := getLocationList(httpClient, api)

	if len(locations) > 1 {
		locations = filterLocations(city, locations)
	}

	index := promptLocationChoice(locations)
	chosenLocation := locations[index]

	return extractCoordinates(chosenLocation.(map[string]interface{}))
}
func promptCity() string {
	fmt.Print("Enter city: ")
	city, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	return city
}
func promptCountryCode() string {
	fmt.Print("Enter Country Code (2 Letters): ")
	countryCode, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	countryCode = strings.TrimRight(countryCode, "\r\n")
	if len(countryCode) != 2 {
		fmt.Println(len(countryCode))
		fmt.Printf("Invalid country code!")
		return promptCountryCode()
	} else {
		return countryCode
	}
}
