package server

import (
	"net/http"

	"github.com/unrolled/render"
)

func apiGetHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Get string }{"GET called"})
	}
}

func apiPostHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Post string }{"POST called"})
	}
}
