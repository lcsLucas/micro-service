package main

import (
	"context"
	"testing"

	proto "github.com/lcslucas/micro-service/proto/usuarios"

	"google.golang.org/grpc"
)

func TestService(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}

	defer conn.Close()

	c := proto.NewServiceGetUserClient(conn)

	req := proto.GetUserRequest{
		Id: 2,
	}

	response, err := c.GetUser(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método GetUser: %s", err)
	}

	if response.Err != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Err)
	}

}
