package main

import (
	"log"
	"net/http"
	"wsw/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RealIP)

	router.Post("/graphql", graphqlHandler())

	log.Fatal(http.ListenAndServe(":8000", router))
}

func graphqlHandler() http.HandlerFunc {
	scheme := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})
	handler := handler.NewDefaultServer(scheme)
	return func(w http.ResponseWriter, r *http.Request) { handler.ServeHTTP(w, r) }
}
