package userprompt

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Zipprompt struct {
}

func (z Zipprompt) GetCoordinates(httpClient *http.Client, apiKey string) (float64, float64, error) {

	zipCode := promptZipCode()
	countryCode := promptCountryCode()

	api := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/zip?zip=%v,%v&appid=%v", zipCode, countryCode, apiKey)

	location, err := getLocation(httpClient, api)

	if err != nil {
		return 0.0, 0.0, err
	}

	lat, long := extractCoordinates(location)

	return lat, long, err
}
func promptZipCode() string {
	fmt.Print("Enter zip code: ")
	zip, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(zip)
}
func getLocation(httpClient *http.Client, apiUrl string) (map[string]interface{}, error){
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

	locationData := x.(map[string]interface{})

	if locationData["message"] != nil {
		return nil, errors.New("AREA NOT FOUND")
	}
	return locationData, nil
}
