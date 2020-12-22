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
	"github.com/lcslucas/micro-service/pedidos"
	proto "github.com/lcslucas/micro-service/proto/pedidos"

	"google.golang.org/grpc"
)

var DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "", "localhost", "3306", "pedidos")

type ServerGRPC struct{}

func (s *ServerGRPC) GetProduto(ctx context.Context, req *proto.GetProdutoRequest) (*proto.GetProdutoResponse, error) {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "pedidos",
			"hour", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	db, err := sql.Open("mysql", DBURL)
	if err != nil {
		level.Error(logger).Log("exit", err)

		return &proto.GetProdutoResponse{
			Pedido: &proto.Pedido{},
			Err:    err.Error(),
		}, nil
	}

	var srv pedidos.Service
	{
		repository := pedidos.NewRepository(db, logger)
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

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "pedidos",
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

	grpcListener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8082))
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}

	s := ServerGRPC{}

	go func() {
		grpcServer := grpc.NewServer()
		proto.RegisterServiceGetProdutoServer(grpcServer, &s)
		if err := grpcServer.Serve(grpcListener); err != nil {
			logger.Log("transport", "gRPC", "during", "listen", "err", "Fatal to serve port 8082:", err)
		}
	}()

	level.Error(logger).Log("exit", <-errs)

}
