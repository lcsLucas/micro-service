package routes

import (
	"net/http"

	"github.com/lcslucas/micro-service/api/handlers"
)

var homeRoutes = []Route{
	Route{
		Uri:     "/",
		Method:  http.MethodGet,
		Handler: handlers.HomeHandler,
	},
}
