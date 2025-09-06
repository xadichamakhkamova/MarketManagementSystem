package models

type ProductSalesReport struct {
	UnitsSold  int32 `json:"units_sold"`
	TotalSales int64 `json:"total_sales"`
	CostPrice  int64 `json:"cost_price"`
	NetProfit  int64 `json:"net_profit"`
}

type ProductSalesUpdateRequest struct {
	CostPrice    int64  `json:"cost_price"`
	SellingPrice int64  `json:"selling_price"`
	ProductId    string `json:"product_id"`
	ProductColor string `json:"product_color"`
}
