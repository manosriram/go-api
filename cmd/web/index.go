package main

import (
	"github.com/manosriram/go-api/pkg/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/home", handlers.Home);

	_ = http.ListenAndServe(":8000", nil);
}