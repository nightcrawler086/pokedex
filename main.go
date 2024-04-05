package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", GetIndex).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", r))

}
