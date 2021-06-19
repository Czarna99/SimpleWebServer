package main

import (
	app "Attractions/SimpleWebServer/application"
	"Attractions/SimpleWebServer/data"
)

var Attraction []data.Attraction

func main() {
	app.HandleRequests()
	Attraction = []data.Attraction{
		data.Attraction{Id: 1, Name: "Eiffel Tower", City: "Paris", Country: "France"},
		data.Attraction{Id: 2, Name: "Tower", City: "London", Country: "Great Britan"},
	}

}
