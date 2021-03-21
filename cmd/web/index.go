package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/manosriram/go-api/pkg/handlers"
	"github.com/manosriram/go-api/pkg/db"
)

func main() {
	r := mux.NewRouter()

	db.Initialize();

	r.HandleFunc("/", handlers.CreateUser).Methods("GET")
	r.HandleFunc("/append", handlers.AddUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
