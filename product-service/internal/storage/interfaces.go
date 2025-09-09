package storage

import (
	pb "product-service/genproto/productpb"
)

type IProductRepository interface {
	CreateProduct(req *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
	GetProductById(req *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error)
	GetProductByFilter(req *pb.GetProductByFilterRequest) (*pb.GetProductByFilterResponse, error)
	UpdateStock(req *pb.UpdateStockRequest) (*pb.UpdateStockResponse, error,)
	UpdateProduct(req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error)
	DeleteProduct(req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error)
}
