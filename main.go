package main

import (
	"fmt"
	"github.com/beer-xplorer-api/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleIndex)
	router.HandleFunc("/beers", controllers.GetBeers)
	router.HandleFunc("/beers/{beerId}", controllers.GetBeer)
	http.ListenAndServe(":"+port, router)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the beer-xplorer-api")
}
