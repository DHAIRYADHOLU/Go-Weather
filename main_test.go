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

	// Capture the printed output
	output := captureOutput(func() {
		printWeatherData(getWeatherData())
	})

	// Check against the expected output
	expectedOutput := "Main: Clear\nDescription: Clear sky\nTemperature: 20.50°C\n"
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
