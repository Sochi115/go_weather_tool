package userprompt

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

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
		countryCode = strings.TrimSpace(countryCode)
		return strings.TrimRight(countryCode, "\r\n")
	}
}
