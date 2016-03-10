package server

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
)

func isAuthorized(formatter *render.Render) negroni.HandlerFunc {
	apikey := os.Getenv(APIKey)
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		providedKey := r.Header.Get(APIKey)
		if providedKey == "" || providedKey != apikey {
			formatter.JSON(w, http.StatusUnauthorized, struct{ Error string }{"Unauthorized."})
		} else {
			next(w, r)
		}
	}
}
