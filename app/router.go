package app

import (
	"net/http"

	"wsw/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
)

func newRouter(midlewares Middlewares) *chi.Mux {
	gqlHandler := createGQLHandler()

	router := chi.NewRouter()
	router.Use(midlewares.List...)

	router.Options("/graphql", gqlHandler)
	router.Post("/graphql", gqlHandler)

	router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("server panic")
	})

	return router
}

func createGQLHandler() http.HandlerFunc {
	scheme := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})
	server := handler.NewDefaultServer(scheme)
	return func(w http.ResponseWriter, r *http.Request) { server.ServeHTTP(w, r) }
}
