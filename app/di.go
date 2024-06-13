package app

import (
	"context"
	"net/http"
	"strings"
	"wsw/backend/ent"
	"wsw/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golobby/container/v3"
	"github.com/rs/cors"
)

func initDi(config Config, appContext context.Context) {
	container.Singleton(func() context.Context { return appContext })
	container.Singleton(func() Config { return config })
	container.Singleton(func(config Config) (*ent.Client, error) {
		client, err := ent.Open("postgres", createConnectionURI(config.Postgres))
		if err != nil {
			return nil, err
		}
		if err := client.Schema.Create(appContext); err != nil {
			return nil, err
		}

		return client, nil
	})
	container.Singleton(func() cors.Options {
		return cors.Options{
			AllowedOrigins:     []string{"*"},
			AllowCredentials:   true,
			AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
			AllowedHeaders:     []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept", "Authorization"},
			OptionsPassthrough: true,
			Debug:              true,
		}
	})

	container.Singleton(func(corsOptions cors.Options) *chi.Mux {
		gqlHandler := createGQLHandler()

		router := chi.NewRouter()
		router.Use(middleware.Logger, middleware.Recoverer, middleware.RealIP, cors.New(corsOptions).Handler)

		router.Options("/graphql", gqlHandler)
		router.Post("/graphql", gqlHandler)
		return router
	})

	container.Singleton(func(router *chi.Mux) App { return appImpl{router} })
}

func createGQLHandler() http.HandlerFunc {
	scheme := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})
	handler := handler.NewDefaultServer(scheme)
	return func(w http.ResponseWriter, r *http.Request) { handler.ServeHTTP(w, r) }
}

func createConnectionURI(config Postgres) string {
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
