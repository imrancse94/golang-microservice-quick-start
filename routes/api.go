package routes

import (
	"github.com/gorilla/mux"
	"go.quick.start/app/requests"
	"go.quick.start/middleware"
	"go.quick.start/register"
)

var AppRouter = register.HTTPRouter{
	Route: []register.Route{
		{
			Name:        "token",
			Path:        "/token",
			Action:      "UserController@Login",
			Method:      "POST",
			Validation:  &requests.Credential{},
			Description: "Test user authentication",
		},
	},
	Groups: []register.Group{
		{
			Name:       "admin",
			Prefix:     "/user",
			Middleware: []register.Middleware{middleware.Auth()},
			Routes: []register.Route{
				{
					Name:        "home",
					Path:        "/home",
					Action:      "UserController@Index",
					Method:      "GET",
					Description: "Main route",
				},
				{
					Name:        "users",
					Path:        "/users",
					Action:      "UserController@RegisterUser",
					Method:      "POST",
					Description: "Insert new user",
				},
				{
					Name:        "auth-data",
					Path:        "/auth-data",
					Action:      "UserController@AuthData",
					Method:      "GET",
					Description: "Test user authentication",
				},
				{
					Name:        "roles",
					Path:        "/roles",
					Action:      "RoleController@GetRole",
					Method:      "GET",
					Description: "Test user authentication",
				},
			},
		},
	},
}

var Router = []register.HTTPRouter{
	AppRouter,
}

func Setup() *mux.Router {
	r := WebRouter(Router)
	return r
}
