package main

import "fmt"

func getUserCity() string {
	var city string

	fmt.Print("Enter city: ")
	fmt.Scanln(&city)

	return city
}