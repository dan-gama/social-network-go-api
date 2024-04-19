package router

import (
	"sn-api/src/router/routes"

	"github.com/gorilla/mux"
)

// NewRouter vai retornar um router com as rotas configuradas
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	return routes.Setup(router)
}
