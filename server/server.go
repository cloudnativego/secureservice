package server

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

//NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()

	// Public Routes
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)

	// Protected API Routes
	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/api/get", apiGetHandler(formatter)).Methods("GET")
	apiRouter.HandleFunc("/api/post", apiGetHandler(formatter)).Methods("POST")

	router.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(isAuthorized(formatter)),
		negroni.Wrap(apiRouter),
	))

	n.UseHandler(router)
	return n
}
