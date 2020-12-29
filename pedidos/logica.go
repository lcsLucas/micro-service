package pedidos

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	proto_usu "github.com/lcslucas/micro-service/proto/usuarios"
	"github.com/lcslucas/micro-service/usuarios"

	logkit "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"google.golang.org/grpc"
)

type service struct {
	repository Repository
	logger     logkit.Logger
}

func (s service) Get(ctx context.Context, id int64) (Pedido, error) {
	logger := logkit.With(s.logger, "method", "Get")

	err := godotenv.Load("../../.env")

	if err != nil {
		return Pedido{}, err
	}

	p, err := s.repository.Get(ctx, id)
	if err != nil {
		level.Error(logger).Log("error", err)
		return Pedido{}, err
	}

	/*Recupera as informações do usuario do pedido*/
	var conn *grpc.ClientConn

	portGrpcUsuarios := fmt.Sprintf("%s:%s", os.Getenv("USU_GRPC_HOST"), os.Getenv("USU_GRPC_PORT"))

	conn, err = grpc.Dial(portGrpcUsuarios, grpc.WithInsecure())
	if err != nil {
		level.Error(logger).Log("Não foi possível conectar: %s", err)
		p.Usu = usuarios.Usuario{
			Err: "Erro inesperado na comunicaçao com o serviço de usuário",
		}
		return p, nil
	}

	defer conn.Close()

	c := proto_usu.NewServiceGetUserClient(conn)

	req := proto_usu.GetUserRequest{
		Id: int64(p.Idusuarios),
	}

	response, err := c.GetUser(context.Background(), &req)
	if err != nil {
		fmt.Printf("Não foi possível chamar o método GetUser: %s", err)
		p.Usu = usuarios.Usuario{
			Err: "Erro inesperado na comunicaçao com o serviço de usuário",
		}
		return p, nil
	}

	if response.Err != "" {
		fmt.Printf("Erro no retorno do serviço de usuário: %s", response.Err)
		p.Usu = usuarios.Usuario{
			Err: "Erro inesperado na comunicaçao com o serviço de usuário",
		}
		return p, nil
	}

	p.Usu = usuarios.Usuario{
		ID:    int(response.Usuario.Id),
		Nome:  response.Usuario.Nome,
		Email: response.Usuario.Email,
	}

	/*Recupera as informações do usuario do pedido*/

	logger.Log("Get Pedido", id)
	return p, nil
}

func NewService(rep Repository, logger logkit.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}
