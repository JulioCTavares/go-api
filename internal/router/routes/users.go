package routes

import (
	"api/internal/controllers"

	"github.com/gorilla/mux"
)

var userRouter = []Route{
	{Path: "/users", Method: "GET", Handler: controllers.GetUsers, IsPrivate: false},
	{Path: "/users/{id}", Method: "GET", Handler: controllers.GetUser, IsPrivate: false},
	{Path: "/users", Method: "POST", Handler: controllers.CreateUser, IsPrivate: false},
	{Path: "/users/{id}", Method: "PUT", Handler: controllers.UpdateUser, IsPrivate: false},
	{Path: "/users/{id}", Method: "DELETE", Handler: controllers.DeleteUser, IsPrivate: false},
}

func HandleUserRoutes(r *mux.Router) *mux.Router {

	for _, route := range userRouter {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	return r
}
