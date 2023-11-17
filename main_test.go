package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetWeatherData(t *testing.T) {
	// Create a fake server to mock the OpenWeatherMap API response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a sample OpenWeatherMap API response
		response := `{"main": {"temp": 20.5}, "weather": [{"main": "Clear", "description": "Clear sky"}]}`
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Override the OpenWeatherMapURL with the fake server's URL
	OpenWeatherMapURL = server.URL

	// Call the function with the fake server's URL
	weatherData := getWeatherData()

	// Check against the expected values
	expectedTemp := 20.5
	if weatherData.Main.Temp != expectedTemp {
		t.Errorf("Expected temperature: %f, but got: %f", expectedTemp, weatherData.Main.Temp)
	}

	expectedMain := "Clear"
	if weatherData.Weather[0].Main != expectedMain {
		t.Errorf("Expected main weather: %s, but got: %s", expectedMain, weatherData.Weather[0].Main)
	}

	expectedDescription := "Clear sky"
	if weatherData.Weather[0].Description != expectedDescription {
		t.Errorf("Expected weather description: %s, but got: %s", expectedDescription, weatherData.Weather[0].Description)
	}
}

func TestPrintWeatherData(t *testing.T) {
	// Create a WeatherData struct with sample data
	weatherData := &WeatherData{
		Main: struct {
			Temp float64 `json:"temp"`
		}{
			Temp: 25.0,
		},
		Weather: []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
		}{
			{
				Main:        "Cloudy",
				Description: "Partly cloudy",
			},
		},
	}

	// Capture the printed output
	output := captureOutput(func() {
		printWeatherData(weatherData)
	})

	// Check against the expected output
	expectedOutput := "Main: Cloudy\nDescription: Partly cloudy\nTemperature: 25.00°C\n"
	if output != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, output)
	}
}

// captureOutput captures the standard output of a function and returns it as a string
func captureOutput(f func()) string {
	// Redirect stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function
	f()

	// Reset stdout
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout

	return string(out)
}
