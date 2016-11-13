package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beer-xplorer-api/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func readBeersList() []models.Beer {
	jsonBeers, err := ioutil.ReadFile("data/beers.json")
	if err != nil {
		fmt.Println("Error while getting beers list", err)
	}
	var beers []models.Beer
	json.Unmarshal([]byte(jsonBeers), &beers)
	return beers
}

func GetBeers(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi((r.FormValue("limit")))
	if err != nil || limit == 0 || limit > 25 {
		limit = 25
	}
	json.NewEncoder(w).Encode(readBeersList()[:limit])
}

func GetBeer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	beerId := params["beerId"]
	beers := readBeersList()
	var beer models.Beer
	for _, b := range beers {
		if b.Id == beerId {
			beer = b
		}
	}
	json.NewEncoder(w).Encode(beer)
}
