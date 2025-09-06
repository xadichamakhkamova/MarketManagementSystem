package handler

import (
	pb "api-gateway/genproto/productpb"
	"api-gateway/logger"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/minio/minio-go/v7"
)

const constanta = "allDataFromDB"

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

func (h *HandlerST) GetProductByIdSocket(c *gin.Context) {

	req := pb.GetProductByIdRequest{}
	req.Id = c.Param("id")

	for {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(1)
			return
		}

		for {
			resp, err := h.service.GetProductById(c, &req)
			if err != nil {
				fmt.Println(2)
				conn.Close()
				return

			}

			if err := conn.WriteJSON(resp); err != nil {
				fmt.Println(3)
				conn.Close()
				return
			}

			select {
			case <-time.After(1):
				continue
			case <-c.Done():
				conn.Close()
				return
			}
		}
	}

}

// @Router /products [post]
// @Summary CREATE PRODUCTS
// @Security BearerAuth
// @Description This method creates products
// @Tags PRODUCTS
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Product Image"
// @Param product formData string true "Product JSON Data"
// @Success 200 {object} models.CreateProductResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) CreateProduct(c *gin.Context) {
	// 1. Faylni olish
	file, err := c.FormFile("file")
	if err != nil {
		logger.Error("CreateProduct: Form file error - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 2. Mahsulot ma'lumotlarini formadagi fielddan olish
	productData := c.PostForm("product")
	if productData == "" {
		logger.Error("CreateProduct: Post form error")
		c.JSON(400, gin.H{"error": "product data is required"})
		return
	}

	// 3. Formadagi JSON satrini parse qilish
	var req pb.CreateProductRequest
	if err := json.Unmarshal([]byte(productData), &req); err != nil {
		logger.Error("CreateProduct: Bind json error - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 4. Faylni vaqtincha saqlash
	path := filepath.Join("./media", file.Filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		logger.Error("CreateProduct: SaveUploadedFile error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer os.Remove(path) // Yuklangan faylni o'chirish

	// 5. Bucket yaratish yoki tekshirish
	bucketName := "products"
	err = h.minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := h.minioClient.BucketExists(context.Background(), bucketName)
		if errBucketExists != nil || !exists {
			logger.Error("CreateProduct: MakeBucket error - ", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	// 6. Faylni MinIO'ga yuklash
	newFileName := uuid.NewString() + filepath.Ext(file.Filename)
	uploadInfo, err := h.minioClient.FPutObject(
		context.Background(),
		bucketName,
		newFileName,
		path,
		minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")},
	)
	if err != nil {
		logger.Error("CreateProduct: FPutObject error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 7. Fayl URL manzilini yaratish
	minioUrl := fmt.Sprintf("http://%s/%s/%s", h.addr, bucketName, uploadInfo.Key)
	req.ImageUrl = minioUrl

	// 8. Mahsulotni ma'lumotlar bazasiga qo'shish
	resp, err := h.service.CreateProduct(context.Background(), &req)
	if err != nil {
		logger.Error("CreateProduct: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 9. Javobni qaytarish
	logger.Info("CreateProduct: Successfully ", resp.Product.Id)
	c.JSON(200, resp)
}

// @Router /products/{id} [get]
// @Summary GET PRODUCT BY ID
// @Security  		BearerAuth
// @Description This method gets product by id
// @Tags PRODUCTS
// @Accept json
// @Produce json
// @Param id path string true "Product Id"
// @Success 200 {object} models.GetProductByIdResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetProductById(c *gin.Context) {

	req := pb.GetProductByIdRequest{}
	req.Id = c.Param("id")
	resp, err := h.service.GetProductById(context.Background(), &req)
	if err != nil {
		logger.Error("GetProductById: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("GetProductById: Successfully ", resp.Product.BagId)
	c.JSON(200, resp)
}

// @Router /products	[get]
// @Summary GET PRODUCTS LIST
// @Security  		BearerAuth
// @Description This method gets products list by filter
// @Tags PRODUCTS
// @Accept json
// @Produce json
// @Param search query string false "Product search"
// @Success 200 {object} models.GetProductByFilterResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetProductByFilter(c *gin.Context) {

	req := pb.GetProductByFilterRequest{}
	search := c.Query("search")

	unknown := strings.TrimSpace(search)
	if unknown == "" {
		req.Search = constanta
	}
	req.Search = unknown
	resp, err := h.service.GetProductByFilter(context.Background(), &req)
	if err != nil {
		logger.Error("GetProductByFilter: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("GetProductByFilter: Successfully ", len(resp.Products))
	c.JSON(200, resp)
}

// @Router /products [put]
// @Summary UPDATE STOCK
// @Security  		BearerAuth
// @Description This method updates stock
// @Tags PRODUCTS
// @Accept json
// @Produce json
// @Param product body models.UpdateStockRequest true "Product"
// @Success 200 {object} models.UpdateStockResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) UpdateStock(c *gin.Context) {

	req := pb.UpdateStockRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("UpdateStock: Error binding json - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.service.UpdateStock(context.Background(), &req)
	if err != nil {
		logger.Error("UpdateStock: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("UpdateStock: Successfully ")
	c.JSON(200, resp)
}

// @Router /products/{id} [put]
// @Summary UPDATE PRODUCTS
// @Security  		BearerAuth
// @Description This method updates products
// @Tags PRODUCTS
// @Accept json
// @Produce json
// @Param id path string true "Product id"
// @Param product body models.UpdateProductRequest true "Product"
// @Success 200 {object} models.UpdateProductResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) UpdateProduct(c *gin.Context) {

	req := pb.UpdateProductRequest{}
	req.Id = c.Param("id")
	if err := c.BindJSON(&req); err != nil {
		logger.Error("UpdateProduct: JSON binding error - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.UpdateProduct(context.Background(), &req)
	if err != nil {
		logger.Error("UpdateProduct: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("UpdateProduct: Successfully ")
	c.JSON(200, resp)
}

// @Router /products/{id} [delete]
// @Summary DELETE PRODUCT
// @Security  		BearerAuth
// @Description This method deletes products
// @Tags PRODUCTS
// @Accept json
// @Produce json
// @Param id path string true "Product Id"
// @Success 200 {object} models.DeleteProductResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) DeleteProduct(c *gin.Context) {

	req := pb.DeleteProductRequest{}
	req.Id = c.Param("id")
	resp, err := h.service.DeleteProduct(context.Background(), &req)
	if err != nil {
		logger.Error("DeleteProduct: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("DeleteProduct: Successfully ", resp.Message)
	c.JSON(200, resp)
}
