package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequests(t *testing.T) {
	// fakeServer := httptest.NewServer(http.HandlerFunc(func(...) {
	// 	w.Write([]byte{})
	// }))

	// fakeServer.URL

	router := newRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	response, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 response, got %d", response.StatusCode)
	}

	response, err = http.Get(server.URL + "/nyc")
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 response, got %d", response.StatusCode)
	}

	var conditions map[string]float64

	dec := json.NewDecoder(response.Body)
	if err := dec.Decode(&conditions); err != nil {
		t.Errorf("Expected json with temperature values, got %s", err)
	}

	feelsLike, ok := conditions["feels_like"]

	if !ok {
		t.Errorf("Expected feels_like key")
	}

	if feelsLike == 0 {
		t.Errorf("Expected feels_like value != 0, got %f", feelsLike)
	}
}
