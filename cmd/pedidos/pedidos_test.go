package main

import (
	"context"
	"testing"

	proto "github.com/lcslucas/micro-service/proto/pedidos"

	"google.golang.org/grpc"
)

func TestService(t *testing.T) {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar ao servidor: %s", err)
	}

	defer conn.Close()

	c := proto.NewServiceGetProdutoClient(conn)

	req := proto.GetProdutoRequest{
		Id: 1,
	}

	response, err := c.GetProduto(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método GetProduto: %s", err)
	}

	if response.Err != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Err)
	}

}
