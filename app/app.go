package app

import (
	"net/http"
	"wsw/backend/lib/utils"

	"github.com/go-chi/chi/v5"
)

func Start(router chi.Router, config Config) {
	utils.F("Error: %v", http.ListenAndServe(":8000", router))
}
