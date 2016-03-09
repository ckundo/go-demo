package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func newRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Home)

	return router
}

func main() {
	router := newRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
