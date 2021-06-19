package app

import (
	"Attractions/SimpleWebServer/data"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var addr = ":8080"

func HandleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", data.HomePage)
	r.HandleFunc("/attractions", data.ReturnAllAtractions)
	r.HandleFunc("/attraction", data.CreateNewAttraction).Methods("POST")
	r.HandleFunc("/attraction/{id}", data.ReturnSingleAttraction)
	r.HandleFunc("/attraction/{id}", data.UpdateAttraction).Methods("PUT")

	log.Println("Application is running on", addr, "...")
	log.Println("Using MUX Routers")
	log.Fatal(http.ListenAndServe(addr, r))
}
