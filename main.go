package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DHAIRYADHOLU/Go-Weather/weather"
	"github.com/gorilla/mux"
)

// APIKey should be set as an environment variable or passed in securely
const APIKey = "your_openweathermap_api_key"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/weather", GetWeather).Methods("GET")

	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// GetWeather is the handler function for the /weather endpoint
func GetWeather(w http.ResponseWriter, r *http.Request) {
	weatherData, err := weather.GetWeatherData(APIKey)
	if err != nil {
		http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		return
	}

	// Convert the weather data to JSON and send it as the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}
