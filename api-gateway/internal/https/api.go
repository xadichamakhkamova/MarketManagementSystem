package https

import (
	"api-gateway/internal/https/handler"
	"api-gateway/internal/service"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Api-gateway service
// @version 1.0
// @description Api-gateway service
// @host localhost:9000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(service *service.ServiceRepositoryClient, port int, minio *minio.Client, addr string) *http.Server {

	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	r.Use(cors.New(corsConfig))

	apiHandler := handler.NewApiHandler(service, minio, addr)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/products", apiHandler.CreateProduct)
	r.PUT("/products", apiHandler.UpdateStock)
	r.PUT("/products/:id", apiHandler.UpdateProduct)
	r.GET("/products/:id/ws", apiHandler.GetProductByIdSocket)
	r.GET("/products", apiHandler.GetProductByFilter)
	r.DELETE("/products/:id", apiHandler.DeleteProduct)

	r.GET("/dashboard", apiHandler.GetDashboardReport)

	r.POST("/debts", apiHandler.CreateDebt)
	r.PUT("/debts/:id", apiHandler.UpdateDebt)
	r.GET("/debts/:id", apiHandler.GetDebtById)
	r.GET("/debts", apiHandler.GetDebtByFilter)
	r.DELETE("/debts/:id", apiHandler.DeleteDebt)

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	address := fmt.Sprintf(":%d", port)
	srv := &http.Server{
		Addr:      address,
		Handler:   r,
		TLSConfig: tlsConfig,
	}

	return srv
}
