package main

import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequests(t *testing.T) {
	ts := httprouter.New()
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)

	addRoutes(ts)

	ts.ServeHTTP(w, r)

	if(w.Code != http.StatusOK) {
		t.Errorf("Expected 200 response, got %d", w.Code)
	}
}
