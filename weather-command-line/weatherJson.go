package main

import (
	"fmt"
	"time"
)

type WeatherData struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Rain struct {
		OneH float64 `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func printWeatherInfo(wd WeatherData) {
	fmt.Println("====== WEATHER FORECAST ======")
	fmt.Println()
	fmt.Println("LOCATION")
	fmt.Printf("%-15v %v\n", "City:", wd.Name)
	fmt.Printf("%-15v %v\n", "Country:", wd.Sys.Country)
	fmt.Printf("%-5v %v   %-5v %v\n", "Sunrise:", parseUnixToTimeString(int64(wd.Sys.Sunrise)),"Sunset:", parseUnixToTimeString(int64(wd.Sys.Sunset)) )
	fmt.Println()
	fmt.Println("WEATHER")
	fmt.Printf("%-15v %v\n", "Weather:", wd.Weather[0].Main )
	fmt.Printf("%-15v %v\n", "Weather Desc:", wd.Weather[0].Description)
}

func parseUnixToTimeString(unix int64) string {
	// tm := time.Unix(unix, 0).Format("Kitchen")
	tm := time.Unix(unix,0)
	return tm.Local().Format(time.Kitchen)
}