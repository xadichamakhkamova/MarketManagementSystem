package handler

import (
	"api-gateway/internal/service"

	"github.com/minio/minio-go/v7"
)

type HandlerST struct {
	service     *service.ServiceRepositoryClient
	minioClient *minio.Client
	addr        string
}

func NewApiHandler(service *service.ServiceRepositoryClient, minio *minio.Client, addr string) *HandlerST {
	return &HandlerST{
		service:     service,
		minioClient: minio,
		addr:        addr,
	}
}
