package api

import (
	"log"
	"net/http"
	"time"

	"github.com/lcslucas/lojavirtual/api/handlers"
	"github.com/lcslucas/lojavirtual/api/middleware"

	"github.com/gorilla/mux"
)

func Run(addr string) {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.SetMiddlewareJSON(handlers.HandlerHome)).Methods("Get")
	r.HandleFunc("/pedido/{id}", middleware.SetMiddlewareJSON(handlers.HandlerPedido)).Methods("Get")

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
