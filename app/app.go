package app

import (
	"context"
	"net/http"
	"wsw/backend/lib/utils"

	"github.com/go-chi/chi/v5"
	"github.com/golobby/container/v3"
)

type App interface {
	Start()
	Close()
}

type appImpl struct {
	router *chi.Mux
}

// Close implements App.
func (a appImpl) Close() {
	panic("unimplemented")
}

// Start implements App.
func (a appImpl) Start() {
	utils.F("Error: %v", http.ListenAndServe(":8000", a.router))
}

func NewApp() (App, error) {
	initDi(newConfig(), context.Background())
	var application App
	err := container.Resolve(&application)
	return application, err
}
