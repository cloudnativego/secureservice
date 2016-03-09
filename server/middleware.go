package server

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
)

func isAuthorized(formatter *render.Render) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		key := r.Header.Get("API-KEY")
		if key == "" || key != "FLUFFY" {
			formatter.JSON(w, http.StatusUnauthorized, struct{ Error string }{"Unauthorized."})
		} else {
			next(w, r)
		}
	}
}
