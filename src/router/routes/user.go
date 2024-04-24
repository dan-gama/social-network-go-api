package routes

import (
	"net/http"
	"sn-api/src/controllers"
)

var userRoutes = []Route{
	{
		URI:           "/users",
		Method:        http.MethodPost,
		Function:      controllers.UserCreate,
		Authenticated: false,
	},
	{
		URI:           "/users",
		Method:        http.MethodGet,
		Function:      controllers.UserGetAll,
		Authenticated: true,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodGet,
		Function:      controllers.UserGet,
		Authenticated: true,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodPut,
		Function:      controllers.UserUpdate,
		Authenticated: true,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodDelete,
		Function:      controllers.UserDelete,
		Authenticated: true,
	},
	{
		URI:           "/users-by-name/{name}",
		Method:        http.MethodGet,
		Function:      controllers.UserGetByName,
		Authenticated: true,
	},
}
