package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/lcslucas/micro-service/api/routes/router"
)

func Run(addr string) {

	r := router.NewRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	fmt.Printf("Servidor escutando em: %s\n", addr)
	log.Fatal(srv.ListenAndServe())

	/*
		r := mux.NewRouter()
		r.HandleFunc("/", middleware.SetMiddlewareJSON(handlers.HomeHandler)).Methods("Get")
		r.HandleFunc("/pedido/{id}", middleware.SetMiddlewareJSON(handlers.PedidoHandler)).Methods("Get")

		srv := &http.Server{
			Handler: r,
			Addr:    addr,
			// Good practice: enforce timeouts for servers you create!
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		fmt.Printf("Servidor rodando em: %s\n", addr)
		log.Fatal(srv.ListenAndServe())
	*/
}
