package rest

import (
	"net/http"

	"github.com/go-chi/render"
)

type ResolveHandlerFunc func(http.ResponseWriter, *http.Request) (any, error)

func RESTHandle(resolveFunc ResolveHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := resolveFunc(w, r)
		if err != nil {
			render.Render(w, r, ErrRender(err))
		} else {
			render.JSON(w, r, result)
		}
	}
}
