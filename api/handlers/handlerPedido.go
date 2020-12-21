package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lcslucas/lojavirtual/api/controllers/responses"
	proto "github.com/lcslucas/lojavirtual/proto/pedidos"
	"google.golang.org/grpc"
)

func HandlerPedido(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro: não foi possível ler o .ENV -> %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Erro inesperado")
		return
	}

	vars := mux.Vars(r)
	param_id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Parametro não enviado")
		return
	}

	/**/

	portGrpcPedidos := fmt.Sprintf(":%s", os.Getenv("PED_GRPC_PORT"))

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(portGrpcPedidos, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Não foi possível conectar ao servidor: %s", err)
		return
	}

	defer conn.Close()

	c := proto.NewServiceGetProdutoClient(conn)

	req := proto.GetProdutoRequest{
		Id: param_id,
	}

	response, err := c.GetProduto(context.Background(), &req)

	if err != nil {
		fmt.Printf("Não foi possível chamar o método GetProduto: %s", err)
		responses.JSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if response.Err != "" {
		fmt.Printf("Erro recebido do servidor: %s", response.Err)
		responses.JSON(w, http.StatusUnprocessableEntity, response.Err)
		return
	} else {
		responses.JSON(w, http.StatusOK, response.Pedido)
	}

}
