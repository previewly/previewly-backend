package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RealIP)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hi"))
		if err != nil {
			log.Fatal("Could not write to response", err)
		}
	})

	log.Fatal(http.ListenAndServe(":8000", router))
}
