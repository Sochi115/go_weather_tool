package main

func main() {
	city := promptUserCity()
	latitude, longitude := getCityCoordinates(city)
	displayWeatherData(latitude, longitude)
}
