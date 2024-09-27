package router

import (
	"api/internal/router/routes"
	"log"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	routes.HandleUserRoutes(r)

	listRoutes(r)
	return r
}

func listRoutes(router *mux.Router) {
	routes := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		methods, _ := route.GetMethods()
		log.Printf("Métodos: %v - Rota: %s", methods, path)
		return nil
	})
	if routes != nil {
		log.Println("Erro ao listar rotas:", routes)
	}
}
