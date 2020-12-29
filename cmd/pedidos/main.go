package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/lcslucas/micro-service/config"
	database "github.com/lcslucas/micro-service/database/mysql"
	"github.com/lcslucas/micro-service/migrations"
	"github.com/lcslucas/micro-service/pedidos"
	proto "github.com/lcslucas/micro-service/proto/pedidos"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

var conn *gorm.DB
var configDB config.ConfigDB

var logger log.Logger

var host_grpc_ped string
var port_grpc_ped int

type ServerGRPC struct{}

func (s *ServerGRPC) GetProduto(ctx context.Context, req *proto.GetProdutoRequest) (*proto.GetProdutoResponse, error) {
	var srv pedidos.Service
	{
		repository := pedidos.NewRepository(conn, logger)
		srv = pedidos.NewService(repository, logger)
	}

	p, err := srv.Get(ctx, req.Id)

	if err != nil {
		level.Error(logger).Log("exit", err)

		return &proto.GetProdutoResponse{
			Pedido: &proto.Pedido{},
			Err:    err.Error(),
		}, nil
	}

	var proItens []*proto.Itens

	for _, pItem := range p.Itens {
		proItens = append(proItens, &proto.Itens{
			Id:        pItem.ID,
			Descricao: pItem.Descricao,
			Preco:     pItem.Preco,
			Qtde:      pItem.Qtde,
		})
	}

	return &proto.GetProdutoResponse{
		Pedido: &proto.Pedido{
			Id:        p.ID,
			Data:      p.Data,
			Valor:     p.ValorTotal,
			Pagamento: p.Pagamento,
			Itens:     proItens,
			Usuario: &proto.Usuario{
				Id:    int64(p.Usu.ID),
				Nome:  p.Usu.Nome,
				Email: p.Usu.Email,
				Err:   p.Usu.Err,
			},
		},
		Err: "",
	}, nil

}

func inicializeDB() error {
	newConn, err := database.Connect(configDB)
	if err != nil {
		return err
	}

	conn = newConn

	return migrations.ExecMigrationPedidos(conn)
}

func inicializeLogger() {
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"service", "users",
		"hour", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)
}

func inicializeVars() error {
	err := godotenv.Load("../../.env")

	if err != nil {
		return err
	}

	configDB = config.ConfigDB{
		Host:     os.Getenv("PED_DB_HOST"),
		User:     os.Getenv("PED_DB_USER"),
		Password: os.Getenv("PED_DB_PASSWORD"),
		Port:     os.Getenv("PED_DB_PORT"),
		DBName:   os.Getenv("PED_DB_NAME"),
	}

	host_grpc_ped = os.Getenv("PED_GRPC_HOST")
	port_grpc_ped, _ = strconv.Atoi(os.Getenv("PED_GRPC_PORT"))

	return nil
}

func main() {
	var err error

	/* Iniciando Logger */
	inicializeLogger()

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	flag.Parse()
	/* Iniciando Logger */

	/* Inicializando variáveis */
	err = inicializeVars()
	if err != nil {
		level.Error(logger).Log("exit", err)
		return
	}
	/* Inicializando variáveis */

	/* Iniciando conexão com o banco*/
	fmt.Println(configDB)
	err = inicializeDB()
	if err != nil {
		level.Error(logger).Log("exit", err)
		return
	}

	sqlDB, _ := conn.DB()
	defer sqlDB.Close()
	/*Iniciando conexão com o banco*/

	/*Iniciando conexão GRPC com o serviço de usuário*/
	grpcListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host_grpc_ped, port_grpc_ped))
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}

	s := ServerGRPC{}

	go func() {
		grpcServer := grpc.NewServer()
		proto.RegisterServiceGetProdutoServer(grpcServer, &s)
		if err := grpcServer.Serve(grpcListener); err != nil {
			logger.Log("transport", "gRPC", "during", "listen", "err", "Fatal to serve:", host_grpc_ped, port_grpc_ped, ":", err)
		}
	}()
	/*Iniciando conexão GRPC com o serviço de usuário*/

	errs := make(chan error)

	// notifica o programa quando ele é encerrado
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	level.Error(logger).Log("exit", <-errs)

}
