package dashboardservice

import (
	pb "api-gateway/genproto/dashboardpb"
	config "api-gateway/internal/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithDashboardService(cfg config.Config) (*pb.DashboardServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.DashboardService.Host, cfg.DashboardService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pb.NewDashboardServiceClient(conn)
	return &userServiceClient, nil
}