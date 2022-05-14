package routes

import (
	"github.com/gorilla/mux"
	"go.quick.start/register"
)

var AppRouter = register.HTTPRouter{
	Route: []register.Route{
		{
			Name:        "home",
			Path:        "/",
			Action:      "UserController@Index",
			Method:      "GET",
			Description: "Main route",
			Middleware:  []register.Middleware{},
		},
		{
			Name:        "users",
			Path:        "/users",
			Action:      "UserController@RegisterUser",
			Method:      "POST",
			Description: "Insert new user",
			Middleware:  []register.Middleware{},
		},
	},
	Groups: []register.Group{
		{
			Name:   "admin",
			Prefix: "/admin",
			Routes: []register.Route{
				{
					Name:        "test",
					Path:        "/test",
					Action:      "UserController@Login",
					Method:      "POST",
					Description: "Test user authentication",
					Middleware:  []register.Middleware{},
				},
			},
			Middleware: []register.Middleware{},
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
