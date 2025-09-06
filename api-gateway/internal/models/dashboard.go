package models

type GetDashboardReportResponse struct {
	TotalSales     int64 `json:"total_sales"`
	TotalCostPrice int64 `json:"total_cost_price"`
	TotalNetProfit int64 `json:"total_net_profit"`
	TotalUnitsSold int64 `json:"total_units_sold"`
}
