package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}

var OpenWeatherMapURL = "https://api.openweathermap.org/data/2.5/weather?q=Toronto&appid=923a2dc915afc4f2bfcc07956e3fef0b"

func main() {
	http.HandleFunc("/", weatherHandler)

	fmt.Println("Server is running on port 9090")

	http.ListenAndServe(":9090", nil)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	weatherData := getWeatherData()

	// Convert temperature from Kelvin to Celsius and then to an integer
	temperatureCelsius := int(kelvinToCelsius(weatherData.Main.Temp))

	htmlTemplate := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Weather Report</title>
		</head>
		<body>
			<h1>Weather Report</h1>
			<p>Main: {{.Main}}</p>
			<p>Description: {{.Description}}</p>
			<p>Temperature: {{.Temperature}}°C</p>
		</body>
		</html>
	`

	tmpl, err := template.New("weather").Parse(htmlTemplate)
	if err != nil {
		fmt.Println("Error parsing HTML template:", err)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{
		"Main":        weatherData.Weather[0].Main,
		"Description": weatherData.Weather[0].Description,
		"Temperature": temperatureCelsius,
	})
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}

func getWeatherData() *WeatherData {
	res, err := http.Get(OpenWeatherMapURL)
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

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	return &weatherData
}

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}
