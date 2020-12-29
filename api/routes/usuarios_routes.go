package routes

import (
	"net/http"

	"github.com/lcslucas/micro-service/api/middleware"

	"github.com/lcslucas/micro-service/api/handlers"
)

var usuarioRoutes = []Route{
	Route{
		Uri:     "/usuarios/test",
		Method:  http.MethodGet,
		Handler: middleware.SetMiddlewareJSON(handlers.UsuarioTestHandler),
	},
}
