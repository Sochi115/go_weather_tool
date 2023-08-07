package main

import (
	"fmt"
	"time"
)

func printForecast(wd WeatherData) {
	fmt.Println("====== WEATHER FORECAST ======")
	printLocationInfo(wd)
	printTempInfo(wd)
	printWeatherInfo(wd)
	printLastUpdated(wd)
}

func printLocationInfo(wd WeatherData){
	fmt.Println()
	fmt.Println("LOCATION")
	fmt.Printf("%-15v %v\n", "City:", wd.Name)
	fmt.Printf("%-15v %v\n", "Country:", wd.Sys.Country)
	fmt.Printf("%-5v %v   %-5v %v\n", "Sunrise:", parseUnixToTimeString(int64(wd.Sys.Sunrise)), "Sunset:", parseUnixToTimeString(int64(wd.Sys.Sunset)))
}
func printTempInfo(wd WeatherData) {
	fmt.Println()
	fmt.Println("TEMPERATURE")
	fmt.Printf("%-15v %v - %v (celsius)\n", "Temp range:", wd.Main.TempMin, wd.Main.TempMax)
	fmt.Printf("%-15v %v%%\n", "Humidity:", wd.Main.Humidity)
	fmt.Printf("%-15v %v (hPa)\n", "Pressure:", wd.Main.Pressure)



}
func printWeatherInfo(wd WeatherData){
	fmt.Println()
	fmt.Println("WEATHER")
	fmt.Printf("%-15v %v\n", "Weather:", wd.Weather[0].Main)
	fmt.Printf("%-15v %v\n", "Weather Desc:", wd.Weather[0].Description)

}

func printLastUpdated(wd WeatherData){
	fmt.Println()
	fmt.Printf("%-15v %v\n", "Last Updated:", parseUnixToTimeString(int64(wd.Dt)))
}
func parseUnixToTimeString(unix int64) string {
	// tm := time.Unix(unix, 0).Format("Kitchen")
	tm := time.Unix(unix, 0)
	return tm.Local().Format(time.Kitchen)
}