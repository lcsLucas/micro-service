package main

import (
	"context"
	"log"

	proto "github.com/lcslucas/lojavirtual/proto/usuarios"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Não foi possível conectar: %s", err)
	}

	defer conn.Close()

	c := proto.NewServiceGetUserClient(conn)

	req := proto.GetUserRequest{
		Id: 1,
	}

	response, err := c.GetUser(context.Background(), &req)
	if err != nil {
		log.Fatalf("Não foi possível chamar o método GetUser: %s", err)
	}

	if response.Err != "" {
		log.Printf("Erro recebido do servidor: %s", response.Err)
	} else {
		log.Printf("Response from server: ID: %v, Nome: %v, Email: %v", response.Usuario.Id, response.Usuario.Nome, response.Usuario.Email)
	}

}
