package routes

import (
	"api/internal/controllers"
	"api/internal/db"
	"api/internal/repositories"
	"api/internal/services"

	"github.com/gorilla/mux"
)

var dbConn, _ = db.ConnectDB()

var userRepo = repositories.NewUserRepository(dbConn)
var userService = services.NewUserService(userRepo)
var userController = controllers.NewUserController(userService)

var userRouter = []Route{
	{Path: "/users", Method: "GET", Handler: controllers.GetUsers, IsPrivate: false},
	{Path: "/users/{id}", Method: "GET", Handler: controllers.GetUser, IsPrivate: false},
	{Path: "/users", Method: "POST", Handler: userController.CreateUser, IsPrivate: false},
	{Path: "/users/{id}", Method: "PUT", Handler: controllers.UpdateUser, IsPrivate: false},
	{Path: "/users/{id}", Method: "DELETE", Handler: controllers.DeleteUser, IsPrivate: false},
}

func HandleUserRoutes(r *mux.Router) *mux.Router {

	for _, route := range userRouter {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	return r
}
