package productservice

import (
	pb "api-gateway/genproto/productpb"
	config "api-gateway/internal/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithProductService(cfg config.Config) (*pb.ProductServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.ProductService.Host, cfg.ProductService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pb.NewProductServiceClient(conn)
	return &userServiceClient, nil
}