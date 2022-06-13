package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.quick.start/api"
	"go.quick.start/app/controller"
	"go.quick.start/application"
	"go.quick.start/register"
	"go.quick.start/tool"
	"net/http"
	"reflect"
	"strings"
)

func WebRouter(routes []register.HTTPRouter) *mux.Router {
	//SingletonIOC = BuildSingletonContainer()
	router := mux.NewRouter()
	//router.Use(gzipMiddleware)

	for _, r := range routes {
		if len(r.Route) > 0 {
			HandleSingleRoute(r.Route, router)
		}

		if len(r.Groups) > 0 {
			HandleGroups(r.Groups, router)
		}

		GiveAccessToPublicFolder(router)
	}

	return router
}

func GiveAccessToPublicFolder(router *mux.Router) {
	publicDirectory := http.Dir(tool.GetDynamicPath("public"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(publicDirectory)))
}

// HandleSingleRoute handles single path parsing.
// This method it's used to parse every single path. If middleware is present a sub-router with will be created
func HandleSingleRoute(routes []register.Route, router *mux.Router) {
	for _, route := range routes {
		hasMiddleware := len(route.Middleware) > 0
		directive := strings.Split(route.Action, "@")
		validation := route.Validation
		if hasMiddleware {
			subRouter := mux.NewRouter()
			subRouter.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
				if err := validateRequest(validation, request); err != nil {
					responseData := api.Response{
						Status:  api.Status("VALIDATION_ERROR"),
						Message: api.StatusMessage("VALIDATION_ERROR"),
						Data:    err,
					}
					api.ErrorResponse(responseData, writer)
					return
				}

				executeControllerDirective(directive, writer, request)
			}).Methods(route.Method)

			subRouter.Use(parseMiddleware(route.Middleware)...)
			router.Handle(route.Path, subRouter).Methods(route.Method)
		} else {
			router.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
				if err := validateRequest(validation, request); err != nil {
					responseData := api.Response{
						Status:  api.Status("VALIDATION_ERROR"),
						Message: api.StatusMessage("VALIDATION_ERROR"),
						Data:    err,
					}
					api.ErrorResponse(responseData, writer)
					return
				}

				executeControllerDirective(directive, writer, request)
			}).Methods(route.Method)
		}
	}
}

// HandleGroups parses route groups.
func HandleGroups(groups []register.Group, router *mux.Router) {
	for _, group := range groups {
		subRouter := router.PathPrefix(group.Prefix).Subrouter()

		for _, route := range group.Routes {
			directive := strings.Split(route.Action, "@")
			validation := route.Validation
			if len(route.Middleware) > 0 {
				nestedRouter := mux.NewRouter()
				fullPath := fmt.Sprintf("%s%s", group.Prefix, route.Path)
				nestedRouter.HandleFunc(fullPath, func(writer http.ResponseWriter, request *http.Request) {
					if err := validateRequest(validation, request); err != nil {
						responseData := api.Response{
							Status:  api.Status("VALIDATION_ERROR"),
							Message: api.StatusMessage("VALIDATION_ERROR"),
							Data:    err,
						}
						api.ErrorResponse(responseData, writer)
						return
					}

					executeControllerDirective(directive, writer, request)
				}).Methods(route.Method)

				nestedRouter.Use(parseMiddleware(route.Middleware)...)
				subRouter.Handle(route.Path, nestedRouter).Methods(route.Method)
			} else {
				subRouter.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
					if err := validateRequest(validation, request); err != nil {
						responseData := api.Response{
							Status:  api.Status("VALIDATION_ERROR"),
							Message: api.StatusMessage("VALIDATION_ERROR"),
							Data:    err,
						}
						api.ErrorResponse(responseData, writer)
						return
					}

					executeControllerDirective(directive, writer, request)
				}).Methods(route.Method)
			}
		}

		subRouter.Use(parseMiddleware(group.Middleware)...)
	}
}

func executeControllerDirective(d []string, w http.ResponseWriter, r *http.Request) {
	cc := GetControllerInterface(d, w, r)
	reflect.ValueOf(cc).MethodByName(d[1]).Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)})
}

func GetControllerInterface(directive []string, w http.ResponseWriter, r *http.Request) interface{} {
	var result interface{}

	controllers := register.ControllerRegister{
		&controller.UserController{},
		&controller.RoleController{},
	}

	for _, contr := range controllers {
		controllerName := reflect.Indirect(reflect.ValueOf(contr)).Type().Name()
		if controllerName == directive[0] {
			application.RegisterBaseController(w, r, &contr)
			result = contr
		}
	}

	return result
}

func parseMiddleware(mwList []register.Middleware) []mux.MiddlewareFunc {
	var midFunc []mux.MiddlewareFunc

	for i := len(mwList) - 1; i > -1; i-- {
		midFunc = append(midFunc, mwList[i].Handle)
	}

	return midFunc
}
func validateRequest(data interface{}, r *http.Request) interface{} {
	if data != nil {
		if err := tool.DecodeJsonRequest(r, &data); err != nil {
			return err
		}

		if err := tool.ValidateRequest(data); err != nil {
			return err
		}
	}

	return nil
}
