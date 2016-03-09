package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequests(t *testing.T) {
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
}
