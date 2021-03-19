package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manosriram/go-api/pkg/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.CreateUser).Methods("GET")
	r.HandleFunc("/append", handlers.AddUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
