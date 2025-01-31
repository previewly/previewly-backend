package rest

import (
	"net/http"

	"github.com/go-chi/render"
)

type ResolveHandlerFunc func(http.ResponseWriter, *http.Request) (interface{}, error)

func RESTHandle(resolveFunc ResolveHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := resolveFunc(w, r)
		if err != nil {
			render.Render(w, r, ErrRender(err))
		}
		render.JSON(w, r, result)
	}
}
