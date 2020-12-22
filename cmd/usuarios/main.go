package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/lcslucas/micro-service/config"
	database "github.com/lcslucas/micro-service/database/mysql"
	"github.com/lcslucas/micro-service/migrations"
	proto "github.com/lcslucas/micro-service/proto/usuarios"
	"github.com/lcslucas/micro-service/usuarios"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var conn *gorm.DB

type ServerGRPC struct{}

func (s *ServerGRPC) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "users",
			"hour", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

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

func inicializeDB() error {
	err := godotenv.Load("../../.env")

	if err != nil {
		return err
	}

	c := config.ConfigDB{
		Host:     os.Getenv("USU_DB_HOST"),
		User:     os.Getenv("USU_DB_USER"),
		Password: os.Getenv("USU_DB_PASSWORD"),
		Port:     os.Getenv("USU_DB_PORT"),
		DBName:   os.Getenv("USU_DB_NAME"),
	}

	newConn, err := database.Connect(c)
	if err != nil {
		return err
	}

	conn = newConn

	return migrations.ExecMigrationUsuarios(conn)
}

func main() {

	/* Iniciando Logger */
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "users",
			"hour", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	flag.Parse()

	/* Iniciando Logger */

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	/* Iniciando conexão com o banco*/
	err := inicializeDB()
	if err != nil {
		level.Error(logger).Log("exit", err)
	}

	sqlDB, _ := conn.DB()
	defer sqlDB.Close()
	/*Iniciando conexão com o banco*/

	/*Iniciando conexão GRPC com o serviço de usuário*/
	grpcListener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8081))
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}

	s := ServerGRPC{}

	go func() {
		grpcServer := grpc.NewServer()
		proto.RegisterServiceGetUserServer(grpcServer, &s)
		if err := grpcServer.Serve(grpcListener); err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", "Fatal to serve gRPC server port 8081")
		}
	}()
	/*Iniciando conexão GRPC com o serviço de usuário*/

	level.Error(logger).Log("exit", <-errs)
}
