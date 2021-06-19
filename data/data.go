package data

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Attraction struct {
	Id      int    `json:"id"`
	Name    string `json:"attraction_name"`
	City    string `json:"city_name"`
	Country string `json:"country_name"`
}

var Attr = []Attraction{
	Attraction{Id: 1, Name: "Eiffel Tower", City: "Paris", Country: "France"},
	Attraction{Id: 2, Name: "Tower", City: "London", Country: "Great Britan"},
}

var Attractions []Attraction

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome on the HomePage!")
}

func ReturnAllAtractions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Attr)
}
