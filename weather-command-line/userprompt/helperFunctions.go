package userprompt

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getLocationList(httpClient *http.Client, apiUrl string) []interface{} {
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

func promptLocationChoice(locations []interface{}) int16 {
	length := len(locations)
	if length <= 1 {
		return 0
	} else {
		fmt.Println("Multiple locations detected, select the correct location:")
		for i, v := range locations {
			area := v.(map[string]interface{})

			state := area["state"]
			country := area["country"]
			fmt.Printf("%v. %v %v \n", i+1, state, country)
		}

		return promptUserIndexChoice(length)
	}
}

func promptUserIndexChoice(length int) int16 {
	fmt.Println()
	fmt.Printf("%-15v", "Enter location index: ")
	choiceString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	choice, err := strconv.Atoi(strings.TrimRight(choiceString, "\r\n"))

	if err != nil {
		log.Fatal(err)
	}

	if choice <= length && choice > 0 {
		return int16(choice - 1)
	} else {
		fmt.Println("Invalid index entered!")
		return promptUserIndexChoice(length)
	}
}
