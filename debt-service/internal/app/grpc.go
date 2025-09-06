package app

import (
	"fmt"
	"net"
	s "debt-service/internal/service"
	"debt-service/logger"

	pb "debt-service/genproto/debtpb"

	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func New(srv s.DebtService, port int) *App {
	grpcServer := grpc.NewServer()
	pb.RegisterDebtServiceServer(grpcServer, &srv)
	return &App{
		gRPCServer: grpcServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	addr := fmt.Sprintf(":%d", a.port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	logger.Info("grpc server started :", a.port)

	if err := a.gRPCServer.Serve(listener); err != nil {
		return err
	}
	return nil
}


func (a *App) Stop() {
	logger.Info("stopping grpc server:", a.port)
	a.gRPCServer.GracefulStop()
}