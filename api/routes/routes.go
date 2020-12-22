package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

/*
	Carrega todas as rotas (Home, Pedidos)
*/
func Load() []Route {
	var allRoutes []Route

	allRoutes = append(allRoutes, homeRoutes...)   // rotas de home
	allRoutes = append(allRoutes, pedidoRoutes...) // rotas de pedido

	fmt.Println(allRoutes)

	return allRoutes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}

	return r
}
