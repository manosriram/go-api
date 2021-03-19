package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Author struct {
	name string `json:"name"`
}

type Book struct {
	id     string  `json:"id"`
	isbn   string  `json:"isbn"`
	author *Author `json:"author"`
}

var Books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(Books)
	fmt.Println(payload)
	w.Write(payload)
}

func main() {
	r := mux.NewRouter()
	Books = append(Books, Book{id: "1", isbn: "1234", author: &Author{name: "Mano"}})
	Books = append(Books, Book{id: "2", isbn: "5234", author: &Author{name: "Sriram"}})
	Books = append(Books, Book{id: "3", isbn: "12313", author: &Author{name: "Mano Sriram"}})

	r.HandleFunc("/", getBooks).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
