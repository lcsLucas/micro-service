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
	database "github.com/lcslucas/micro-service/database/mongodb"
	"github.com/lcslucas/micro-service/migrations"
	proto "github.com/lcslucas/micro-service/proto/usuarios"
	"github.com/lcslucas/micro-service/usuarios"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var conn *mongo.Client
var configDB config.ConfigDB

var logger log.Logger

var host_grpc_usu string
var port_grpc_usu int

type ServerGRPC struct{}

func (s *ServerGRPC) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	var srv usuarios.Service
	{
		repository := usuarios.NewRepository(conn, logger)
		srv = usuarios.NewService(repository, logger)
	}

	u, err := srv.Get(ctx, int(req.Id))
	if err != nil {
		level.Error(logger).Log("exit", err)

		return &proto.GetUserResponse{
			Usuario: &proto.User{},
			Err:     err.Error(),
		}, nil
	}

	return &proto.GetUserResponse{
		Usuario: &proto.User{
			Id:    int64(u.ID),
			Nome:  u.Nome,
			Email: u.Email,
		},
		Err: "",
	}, nil
}

func inicializeDB(ctx context.Context) error {
	newConn, err := database.Connect(ctx, configDB)
	if err != nil {
		fmt.Println("*** Teste de percurso: cmd/usuarios.go:66 ***")
		return err
	}

	conn = newConn
	return migrations.ExecMigrationUsuarios(conn)
}

/*
func inicializeDB() error {
	newConn, err := database.Connect(configDB)
	if err != nil {
		return err
	}

	conn = newConn

	return migrations.ExecMigrationUsuarios(conn)
}
*/

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
		Driver:   os.Getenv("USU_DB_DRIVER"),
		Host:     os.Getenv("USU_DB_HOST"),
		User:     os.Getenv("USU_DB_USER"),
		Password: os.Getenv("USU_DB_PASSWORD"),
		Port:     os.Getenv("USU_DB_PORT"),
		DBName:   os.Getenv("USU_DB_NAME"),
	}

	host_grpc_usu = os.Getenv("USU_GRPC_HOST")
	port_grpc_usu, _ = strconv.Atoi(os.Getenv("USU_GRPC_PORT"))

	return nil
}

func main() {
	var err error

	ctx := context.Background()

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
	}
	/* Inicializando variáveis */

	/* Iniciando conexão com o banco*/
	err = inicializeDB(ctx)
	if err != nil {
		level.Error(logger).Log("exit", err)
	}
	defer conn.Disconnect(ctx)
	/*Iniciando conexão com o banco*/

	/*Iniciando conexão GRPC com o serviço de usuário*/
	grpcListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host_grpc_usu, port_grpc_usu))
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}

	s := ServerGRPC{}

	go func() {
		grpcServer := grpc.NewServer()
		proto.RegisterServiceGetUserServer(grpcServer, &s)
		if err := grpcServer.Serve(grpcListener); err != nil {
			logger.Log("transport", "gRPC", "during", "listen", "err", "Fatal to serve:", host_grpc_usu, port_grpc_usu, ":", err)
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
