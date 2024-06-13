package app

import (
	"context"
	"net/http"
	"os"
	"wsw/backend/lib/utils"

	"github.com/go-chi/chi/v5"
	"github.com/golobby/container/v3"
	"gopkg.in/yaml.v2"
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
	initDi(readConfig(), context.Background())
	var application App
	err := container.Resolve(&application)
	return application, err
}

func readConfig() Config {
	f, err := os.Open("config.yaml")
	if err != nil {
		panic("cannot open config file")
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		utils.D(err, cfg)
		panic("cannot parse config ")
	}
	return cfg
}
