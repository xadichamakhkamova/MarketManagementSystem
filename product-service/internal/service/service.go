package service

import (
	"context"
	"encoding/json"
	pb "product-service/genproto/productpb"
	"product-service/internal/kafka/producer"
	"product-service/internal/storage"
	"product-service/logger"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	productRepo storage.IProductRepository
	producer    producer.IProducerInit
}

func NewProductService(repo storage.IProductRepository, producer producer.IProducerInit) *ProductService {
	return &ProductService{
		productRepo: repo,
		producer:    producer,
	}
}

func (service *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	return service.productRepo.CreateProduct(req)
}

func (service *ProductService) GetProductById(ctx context.Context, req *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error) {
	return service.productRepo.GetProductById(req)
}

func (service *ProductService) GetProductByFilter(ctx context.Context, req *pb.GetProductByFilterRequest) (*pb.GetProductByFilterResponse, error) {
	return service.productRepo.GetProductByFilter(req)
}

func (service *ProductService) UpdateStock(ctx context.Context, req *pb.UpdateStockRequest) (*pb.UpdateStockResponse, error) {

	resp, err := service.productRepo.UpdateStock(req)
	if err != nil {
		return nil, err
	}

	kafkaMessage := map[string]interface{}{
		"cost_price":    req.CostPrice,
		"selling_price": req.SellingPrice,
		"product_id":    req.ProductId,
		"product_color": req.ProductColor,
	}

	message, err := json.Marshal(kafkaMessage)
	if err != nil {
		logger.Error("Failed to serialize message to JSON: ", err)
		return nil, err
	}

	err = service.producer.ProduceMessage(req.ProductId, message)
	if err != nil {
		logger.Error("Failed to produce message to Kafka: ", err)
		return nil, err
	}

	return resp, nil
}

func (service *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	return service.productRepo.UpdateProduct(req)
}

func (service *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	return service.productRepo.DeleteProduct(req)
}
