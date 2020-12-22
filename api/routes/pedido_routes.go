package routes

import (
	"net/http"

	"github.com/lcslucas/micro-service/api/middleware"

	"github.com/lcslucas/micro-service/api/handlers"
)

var pedidoRoutes = []Route{
	Route{
		Uri:     "/pedido/{id}",
		Method:  http.MethodGet,
		Handler: middleware.SetMiddlewareJSON(handlers.PedidoHandler),
	},
}
