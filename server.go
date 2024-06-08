package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
