package handler

import (
	pb "api-gateway/genproto/dashboardpb"
	"api-gateway/logger"
	"context"

	"github.com/gin-gonic/gin"
)

// @Router /dashboard [get]
// @Summary GET DASHBOARD REPORTS
// @Security  		BearerAuth
// @Description This method gets dashboard reports
// @Tags DASHBOARD
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetDashboardReport(c *gin.Context) {

	resp, err := h.service.GetDashboardReport(context.Background(), &pb.GetDashboardReportRequest{})
	if err != nil {
		logger.Error("GetDashboardReport: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("GetDashboardReport: Successfully retrieved report")
	c.JSON(200, gin.H{
		"total_sales": resp.TotalSales,
		"total_cost_price": resp.TotalCostPrice,
		"total_net_profit": resp.TotalNetProfit,
		"total_units_sold": resp.TotalUnitsSold,
	})
}
