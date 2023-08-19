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

func promptZipOrCityChoice() int16{
	printInitialMenu()
	return promptMenuChoice()
}

func printInitialMenu() {
	fmt.Println("======== WEATHER CLI =========")
	fmt.Println()
	fmt.Println("How would you like to enter location?")
	fmt.Println()
	fmt.Println("1. City + Country")
	fmt.Println("2. Zip code + Country")
}

func promptMenuChoice() int16 {
	fmt.Println()
	fmt.Printf("%-15v", "Enter choice number: ")
	choiceString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	choice, err := strconv.Atoi(strings.TrimRight(choiceString, "\r\n"))

	if err != nil {
		log.Fatal(err)
	}

	if choice <= 2 && choice > 0 {
		return int16(choice)
	} else {
		fmt.Println("Invalid number entered!")
		return promptMenuChoice()
	}
}
