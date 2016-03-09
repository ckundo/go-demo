package main

import(
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func addRoutes(router *httprouter.Router) {
	router.GET("/", Home)
}

func main() {
	router := httprouter.New()
	addRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
