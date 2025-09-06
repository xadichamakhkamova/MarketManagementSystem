package models

type Product struct {
	ImageUrl     string           `bson:"image_url" json:"image_url"`
	ID           string           `json:"id" bson:"_id,omitempty"`
	Name         string           `json:"name" bson:"name"`
	UniqueNumber string           `json:"unique_number" bson:"unique_number"`
	BagID        string           `json:"bag_id" bson:"bag_id"`
	Price        int64            `json:"price" bson:"price"`
	Size         string           `json:"size" bson:"size"`
	Colors       map[string]int32 `json:"colors" bson:"colors"`
	Count        int32            `json:"count" bson:"count"`
	Timestamp    []CUDP           `json:"timestamp" bson:"timestamp"`
}

type CreateProductRequest struct {
	Name         string           `json:"name" bson:"name"`
	UniqueNumber string           `json:"unique_number" bson:"unique_number"`
	BagID        string           `json:"bag_id" bson:"bag_id"`
	Price        int64            `json:"price" bson:"price"`
	Size         string           `json:"size" bson:"size"`
	Colors       map[string]int32 `json:"colors" bson:"colors"`
	Count        int32            `json:"count" bson:"count"`
}

type CreateProductResponse struct {
	Product Product `json:"product" bson:"product"`
}

type GetProductByIdResponse struct {
	Product Product `json:"product" bson:"product"`
}

type GetProductByFilterRequest struct {
	Search string `json:"search" bson:"search"`
}

type GetProductByFilterResponse struct {
	Products []Product `json:"products" bson:"products"`
}

type UpdateStockRequest struct {
	CostPrice    int64  `json:"cost_price"`
	SellingPrice int64  `json:"selling_price"`
	ProductID    string `json:"product_id"`
	ProductColor string `json:"product_color"`
}

type UpdateStockResponse struct {
	IsSuccess bool `json:"is_success"`
}

type UpdateProductRequest struct {
	ImageUrl     string           `bson:"image_url" json:"image_url"`
	ID           string           `json:"id" bson:"id"`
	Name         string           `json:"name" bson:"name"`
	UniqueNumber string           `json:"unique_number" bson:"unique_number"`
	BagID        string           `json:"bag_id" bson:"bag_id"`
	Price        int64            `json:"price" bson:"price"`
	Size         string           `json:"size" bson:"size"`
	Colors       map[string]int32 `json:"colors" bson:"colors"`
	Count        int32            `json:"count" bson:"count"`
}

type UpdateProductResponse struct {
	Product Product `json:"product" bson:"product"`
}

type DeleteProductRequest struct {
	ID string `json:"id" bson:"id"`
}

type DeleteProductResponse struct {
	Name    string `json:"name" bson:"name"`
	Message string `json:"message" bson:"message"`
}

type CUDP struct {
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	DeletedAt int32  `json:"deleted_at" bson:"deleted_at"`
}
