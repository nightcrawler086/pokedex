package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", GetIndex).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server is running on port 8080")

}
