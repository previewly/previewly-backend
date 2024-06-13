package main

import (
	"context"
	"net/http"
	"os"
	"wsw/backend/app"
	"wsw/backend/graph"
	"wsw/backend/lib/utils"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golobby/container/v3"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	app.InitDi(readConfig(), context.Background())

	var config app.Config

	err := container.Resolve(&config)
	if err != nil {
		utils.F("Could not resolve config: %v", err)
	}

	corsMiddleware := cors.
		New(cors.Options{
			AllowedOrigins:     []string{"*"},
			AllowCredentials:   true,
			AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
			AllowedHeaders:     []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept", "Authorization"},
			OptionsPassthrough: true,
			Debug:              true,
		}).
		Handler

	gqlHandler := graphqlHandler()

	router := chi.NewRouter()
	router.Use(middleware.Logger, middleware.Recoverer, middleware.RealIP, corsMiddleware)

	router.Options("/graphql", gqlHandler)
	router.Post("/graphql", gqlHandler)

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
