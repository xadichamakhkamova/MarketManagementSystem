package debtservice

import (
	pb "api-gateway/genproto/debtpb"
	config "api-gateway/internal/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithDebtService(cfg config.Config) (*pb.DebtServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.DebtService.Host, cfg.DebtService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pb.NewDebtServiceClient(conn)
	return &userServiceClient, nil
}