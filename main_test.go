package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWeatherHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/weather", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(weatherHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Weather Report"
	if body := rr.Body.String(); body != "" && !strings.Contains(body, expected) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}
