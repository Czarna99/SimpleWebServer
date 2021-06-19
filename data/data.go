package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func CreateNewAttraction(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var a Attraction
	json.Unmarshal(reqBody, &a)
	Attr = append(Attr, a)
	json.NewEncoder(w).Encode(a)
}

func ReturnSingleAttraction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	x := vars["id"]
	key, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Key: %d", key)

	for _, attraction := range Attr {
		if attraction.Id == key {
			json.NewEncoder(w).Encode(attraction)
		}
	}
}

//TODO - is not working. Not updating records.
func UpdateAttraction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	x := vars["id"]
	id, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}
	var updateEvent Attraction
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(reqBody, &updateEvent)
	for i, a := range Attr {
		if a.Id == id {
			a.City = updateEvent.City
			a.Country = updateEvent.Country
			a.Name = updateEvent.Name
			Attr[i] = a
			json.NewEncoder(w).Encode(a)
		}
	}
}

//TODO - fix, is not working. Specified Attraction remains undeleted
func DeleteAttraction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	x := vars["id"]
	id, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}
	for i, attraction := range Attr {
		if attraction.Id == id {
			Attr = append(Attr[:i], Attr[i+1:]...)
		}
	}
}
