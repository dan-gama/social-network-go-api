package routes

import (
	"net/http"
	"sn-api/src/controllers"
)

var loginRoutes = Route{
	URI:           "/login",
	Method:        http.MethodPost,
	Function:      controllers.LoginAuth,
	Authenticated: false,
}
