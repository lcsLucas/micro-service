package router

import (
	"github.com/gorilla/mux"
	"github.com/lcslucas/micro-service/api/routes"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}
