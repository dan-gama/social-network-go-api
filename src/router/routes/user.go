package routes

import (
	"net/http"
	"sn-api/src/controllers"
)

var userRoutes = []Route{
	{
		URI:           "/users",
		Method:        http.MethodPost,
		Function:      controllers.Create,
		Authenticated: false,
	},
	{
		URI:           "/users",
		Method:        http.MethodGet,
		Function:      controllers.GetAll,
		Authenticated: false,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodGet,
		Function:      controllers.Get,
		Authenticated: false,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodPut,
		Function:      controllers.Update,
		Authenticated: false,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodDelete,
		Function:      controllers.Delete,
		Authenticated: false,
	},
	{
		URI:           "/users-by-name/{name}",
		Method:        http.MethodGet,
		Function:      controllers.GetByName,
		Authenticated: false,
	},
}
