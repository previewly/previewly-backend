package main

import (
	"context"
	"net/http"
	"os"
	"strings"
	"wsw/backend/app"
	"wsw/backend/ent"
	"wsw/backend/graph"
	"wsw/backend/lib/utils"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	config := readConfig()
	appContext := context.Background()

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

	client, err := NewDBClient(appContext, config.Postgres)
	if err != nil {
		utils.F("failed opening connection to postgres: %v", err)
	}
	utils.D(client)

	app.Start(router, config)
}

func NewDBClient(appContext context.Context, config app.Postgres) (*ent.Client, error) {
	client, err := ent.Open("postgres", createConnectionURI(config))
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(appContext); err != nil {
		return nil, err
	}

	return client, nil
}

func createConnectionURI(config app.Postgres) string {
	var sb strings.Builder

	optionsMap := map[string]string{
		"host":     config.Host,
		"port":     config.Port,
		"user":     config.User,
		"password": config.Password,
		"dbname":   config.DB,
		"sslmode":  "disable",
	}

	for key, val := range optionsMap {
		sb.WriteString(key + "=" + val + " ")
	}

	return sb.String()
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
