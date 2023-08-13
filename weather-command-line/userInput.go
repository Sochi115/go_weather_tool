package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)

func getUserCity() string {

	fmt.Print("Enter city: ")
	city, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err.Error())
	}

	return city
}

func getLocationIndex(locations []interface{}) int16{

	length := len(locations)
	if (length <= 1) {
		return 0
	} else {
		fmt.Println("Multiple locations detected, select the correct location:")
		for i, v := range locations {
			area := v.(map[string]interface{})

			state := area["state"]
			country := area["country"]
			fmt.Printf("%v. %v %v \n", i+1, state, country)
		}

		return getUserIndexChoice(length)
	}
}

func getUserIndexChoice(length int) int16 {
	fmt.Println()
	fmt.Printf("%-15v", "Enter location index: ")
	choiceString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	choice, err:= strconv.Atoi(choiceString)

	if err != nil {
		fmt.Println(err.Error())
	}

	if choice <= length{
		return int16(choice - 1)
	} else {
		fmt.Println("Invalid index entered!")
		return getUserIndexChoice(length)
	}
}