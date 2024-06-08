package main

import (
	"net/http"
	"os"
	"wsw/backend/app"
	"wsw/backend/graph"
	"wsw/backend/lib/utils"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	config := readConfig()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RealIP)

	router.Post("/graphql", graphqlHandler())

	app.Start(router, config)
}

func graphqlHandler() http.HandlerFunc {
	scheme := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})
	handler := handler.NewDefaultServer(scheme)
	return func(w http.ResponseWriter, r *http.Request) { handler.ServeHTTP(w, r) }
}

func readConfig() app.Config {
	f, err := os.Open("config.yaml")
	if err != nil {
		panic("cannot open config file")
	}
	defer f.Close()

	var cfg app.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		utils.D(err, cfg)
		panic("cannot parse config ")
	}
	return cfg
}
