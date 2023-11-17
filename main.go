package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// WeatherData struct to represent the JSON response from OpenWeatherMap API
type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}

func main() {
	printWeatherData(getWeatherData())
}

func getWeatherData() *WeatherData {
	url := "https://api.openweathermap.org/data/2.5/weather?q=Toronto&appid=923a2dc915afc4f2bfcc07956e3fef0b"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to fetch weather data:", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	// Parse the JSON response into WeatherData struct
	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	return &weatherData
}

func printWeatherData(data *WeatherData) {
	main := data.Weather[0].Main
	description := data.Weather[0].Description
	temperature := data.Main.Temp

	fmt.Printf("Main: %s\nDescription: %s\nTemperature: %.2f°C\n", main, description, temperature)
}
