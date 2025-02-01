package app

import (
	"context"
	"net/http"

	"wsw/backend/graph"
	"wsw/backend/lib/rest"
	"wsw/backend/resolvers/token"
	"wsw/backend/resolvers/url"

	"github.com/go-chi/chi/v5"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func newRouter(midlewares Middlewares) *chi.Mux {
	srv := createGQLServer(createSchema())
	gqlHandler := func(w http.ResponseWriter, r *http.Request) { srv.ServeHTTP(w, r) }

	router := chi.NewRouter()
	router.Use(midlewares.List...)

	router.Options("/graphql", gqlHandler)
	router.Post("/graphql", gqlHandler)

	router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("server panic")
	})

	router.Post("/json/create-token", rest.RESTHandle(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		return token.ResolveCreateToken(r.Context())
	}))

	router.Post("/json/add-url/{url}/token/{token}/", rest.RESTHandle(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		return url.ResolveAddURL(chi.URLParam(r, "token"), chi.URLParam(r, "url"))
	}))

	router.Get("/json/get-preview/token/{token}/?url={url}", rest.RESTHandle(func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		return url.ResolveGetPreview(chi.URLParam(r, "token"), chi.URLParam(r, "url"))
	}))

	return router
}

func createSchema() graphql.ExecutableSchema {
	return graph.NewExecutableSchema(
		graph.Config{Resolvers: &graph.Resolver{}},
	)
}

func createGQLServer(schema graphql.ExecutableSchema) *handler.Server {
	srv := handler.New(schema)

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		return gqlerror.Errorf("Internal server error!")
	})

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	return srv
}
