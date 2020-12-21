package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/lcslucas/lojavirtual/proto/pedidos"

	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Não foi possível conectar ao servidor: %s", err)
	}

	defer conn.Close()

	c := proto.NewServiceGetProdutoClient(conn)

	req := proto.GetProdutoRequest{
		Id: 1,
	}

	response, err := c.GetProduto(context.Background(), &req)
	if err != nil {
		log.Fatalf("Não foi possível chamar o método GetProduto: %s", err)
	}

	if response.Err != "" {
		log.Fatalf("Erro recebido do servidor: %s", response.Err)
	} else {
		fmt.Println(response.Pedido)
		// log.Fatalf(response.Pedido)
	}

}
