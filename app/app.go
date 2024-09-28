package app

import (
	"context"
	"net/http"
	"strconv"

	"wsw/backend/app/config"
	"wsw/backend/lib/utils"

	"github.com/go-chi/chi/v5"
	"github.com/golobby/container/v3"
)

type (
	closer func()
	App    interface {
		Start()
		Closer() closer
	}
)

type appImpl struct {
	router *chi.Mux
	listen config.ListenHost
	closer closer
}

// Close implements App.
func (a appImpl) Closer() closer {
	return a.closer
}

// Start implements App.
func (a appImpl) Start() {
	utils.F("Error: %v", http.ListenAndServe(a.listen.Host+":"+strconv.Itoa(a.listen.Port), a.router))
}

func NewApp() (App, error) {
	initDi(config.NewConfig(), context.Background())
	var application App
	err := container.Resolve(&application)
	return application, err
}
