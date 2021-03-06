package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	proto "github.com/lcslucas/micro-service/proto/usuarios"
	"google.golang.org/grpc"
)

func UsuarioTestHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro: não foi possível ler o .ENV -> %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Erro inesperado")
		return
	}

	portGrpcUsuarios := fmt.Sprintf("%s:%s", os.Getenv("USU_GRPC_HOST"), os.Getenv("USU_GRPC_PORT"))

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(portGrpcUsuarios, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Não foi possível conectar: %s", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Erro inesperado: %s", err)
	}

	defer conn.Close()

	c := proto.NewServiceGetUserClient(conn)

	req := proto.GetUserRequest{
		Id: 1,
	}

	response, err := c.GetUser(ctx, &req)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Não foi possível chamar o método GetUser: %s", err)
	}

	if response.Err != "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Erro recebido do servidor: %s", response.Err)
	}

	w.WriteHeader(http.StatusOK)
	r_json, _ := json.Marshal(response.Usuario)
	fmt.Fprint(w, string(r_json))

}
