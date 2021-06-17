package main

import (
	"github.com/pawel/attractions/data"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"
	http.HandleFunc("/", data.MainPage)
	http.HandleFunc("/poland", data.Poland)
	http.HandleFunc("/france", data.France)

	log.Println("Application is running on localhost:8080....")
	log.Fatal (http.ListenAndServe(addr, nil))
}




