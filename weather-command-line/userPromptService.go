package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func promptUserCity() string {
	fmt.Print("Enter city: ")
	city, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	return city
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

	if choice <= length {
		return int16(choice - 1)
	} else {
		fmt.Println("Invalid index entered!")
		return promptUserIndexChoice(length)
	}
}
