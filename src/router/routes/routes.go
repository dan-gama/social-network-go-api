package routes

import (
	"net/http"
	"sn-api/src/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	URI           string
	Method        string
	Function      func(http.ResponseWriter, *http.Request)
	Authenticated bool
}

func Setup(router *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoutes)

	for _, route := range routes {
		if route.Authenticated {
			router.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
			continue
		} else {
			router.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return router
}
