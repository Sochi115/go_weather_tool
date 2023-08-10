package main

import "fmt"

func getUserCity() string {
	var city string

	fmt.Print("Enter city: ")
	fmt.Scanln(&city)

	return city
}

func getLocationIndex(locations []interface{}) int16{
	fmt.Println("Multiple locations detected, select the correct location:")

	length := len(locations)
	for i, v := range locations {
		area := v.(map[string]interface{})

		state := area["state"]
		country := area["country"]
		fmt.Printf("%v. %v %v \n", i+1, state, country)
	}

	return getUserIndexChoice(length)
}

func getUserIndexChoice(length int) int16 {
	var choice int16;
	fmt.Println()
	fmt.Printf("%-15v", "Enter location index: ")
	fmt.Scanln(&choice)

	if choice <= int16(length){
		return choice -1
	} else {
		fmt.Println("Invalid index entered!")
		return getUserIndexChoice(length)
	}
}