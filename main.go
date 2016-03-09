package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type WeatherResponse struct {
	CurrentObservation struct {
		FeelsLike string `json:"feelslike_f"`
	} `json:"current_observation"`
}

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func Conditions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)

	apiEndpoint := "http://api.wunderground.com/api/<key>/conditions/q/CA/San_Francisco.json"
	response, err := http.Get(apiEndpoint)
	defer response.Body.Close()
	if err != nil {
		log.Printf("Failed to get conditions: %s\n", err)
		return
	}

	var data WeatherResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		panic(err)
	}

	feelsLike, err := strconv.ParseFloat(data.CurrentObservation.FeelsLike, 64)
	if err != nil {
		log.Printf("Failed to convert json from %s\n", err)
	}

	conditions := map[string]float64{
		"feels_like": feelsLike,
	}

	json.NewEncoder(w).Encode(conditions)
}

func newRouter() http.Handler {
	router := httprouter.New()
	router.GET("/", Home)
	router.GET("/nyc", Conditions)

	return router
}

func main() {
	router := newRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
