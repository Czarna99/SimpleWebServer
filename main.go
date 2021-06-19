package main

import (
	app "Attractions/SimpleWebServer/application"
	"Attractions/SimpleWebServer/data"
)

var Attraction []data.Attraction

func main() {
	app.HandleRequests()

}
