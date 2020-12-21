package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/go-sql-driver/mysql"
	proto "github.com/lcslucas/lojavirtual/proto/usuarios"
	"github.com/lcslucas/lojavirtual/usuarios"
	"google.golang.org/grpc"
)

var DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "", "localhost", "3306", "pedidos")

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

	db, err := sql.Open("mysql", DBURL)
	if err != nil {
		level.Error(logger).Log("exit", err)

		return &proto.GetUserResponse{
			Usuario: &proto.User{},
			Err:     err.Error(),
		}, nil
	}

	var srv usuarios.Service
	{
		repository := usuarios.NewRepository(db, logger)
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

func main() {

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

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

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

	level.Error(logger).Log("exit", <-errs)

}
