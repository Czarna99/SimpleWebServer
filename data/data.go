package data

import (
	"fmt"
	"net/http"
)

type City struct {
	Name string
}

type Attraction struct {
	Name string
	City
	Country
}

type Country struct {
	Name string
}
//MainPage - main page of the website
func MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a main page")
}

//Poland - Displays in web browser country, city and attraction in it.
func Poland(w http.ResponseWriter, r *http.Request) {
	var d Attraction
	d.Country.Name = "Poland"
	d.City.Name = "Warsaw"
	d.Name = "Chopin Statue"

	fmt.Fprintf(w, "Country: %s \n", d.Country.Name)
	fmt.Fprintf(w, "City: %s \n", d.City.Name)
	fmt.Fprintf(w, "Attraction: %s \n", d.Name)

}
//France - Displays in web browser country, city and attraction in it.
func France(w http.ResponseWriter, r *http.Request) {
	var d Attraction
	d.Country.Name = "France"
	d.City.Name = "Paris"
	d.Name = "Eiffel Tower"
	fmt.Fprintf(w, "Country: %s \n", d.Country.Name)
	fmt.Fprintf(w, "City: %s \n", d.City.Name)
	fmt.Fprintf(w, "Attraction: %s \n", d.Name)
}

