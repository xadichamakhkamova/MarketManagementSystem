package handler

import (
	pb "api-gateway/genproto/debtpb"
	"api-gateway/logger"
	"context"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Router /debts [post]
// @Summary CREATE DEBTS
// @Security  		BearerAuth
// @Description This method creates debts
// @Tags DEBTS
// @Accept json
// @Produce json
// @Param product body models.CreateDebtReq true "Debt"
// @Success 200 {object} models.DebtResp
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) CreateDebt(c *gin.Context) {

	req := pb.CreateDebtReq{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("CreateDebt: Error binding json - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.service.CreateDebt(context.Background(), &req)
	if err != nil {
		logger.Error("CreateDebt: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("CreateDebt: Successfully ", resp)
	c.JSON(200, resp)
}

// @Router /debts/{id} [put]
// @Summary UPDATE DEBTS
// @Security  		BearerAuth
// @Description This method updates debts
// @Tags DEBTS
// @Accept json
// @Produce json
// @Param id path string true "Debt id"
// @Param product body models.UpdateDebtReq true "Debt"
// @Success 200 {object} models.DebtResp
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) UpdateDebt(c *gin.Context) {

	req := pb.UpdateDebtReq{}
	req.Id = c.Param("id")
	if err := c.BindJSON(&req); err != nil {
		logger.Error("UpdateDebt: Error binding json - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.service.UpdateDebt(context.Background(), &req)
	if err != nil {
		logger.Error("UpdateDebt: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("UpdateDebt: Successfully ", resp)
	c.JSON(200, resp)
}

// @Router /debts/{id} [delete]
// @Summary DELETE DEBTS
// @Security  		BearerAuth
// @Description This method deletes debts
// @Tags DEBTS
// @Accept json
// @Produce json
// @Param id path string true "Debt Id"
// @Success 200 {object} models.DebtResp
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) DeleteDebt(c *gin.Context) {

	req := pb.DeleteDebtReq{}
	req.Id = c.Param("id")
	resp, err := h.service.DeleteDebt(context.Background(), &req)
	if err != nil {
		logger.Error("DeleteDebt: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("DeleteDebt: Successfully ", resp)
	c.JSON(200, resp)
}

// @Router /debts/{id} [get]
// @Summary GET DEBTS BY ID
// @Security  		BearerAuth
// @Description This method gets debt by id
// @Tags DEBTS
// @Accept json
// @Produce json
// @Param id path string true "Debt Id"
// @Success 200 {object} models.DebtResp
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetDebtById(c *gin.Context) {

	req := pb.GetDebtByIdReq{}
	req.Id = c.Param("id")
	resp, err := h.service.GetDebtById(context.Background(), &req)
	if err != nil {
		logger.Error("GetDebtById: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("GetDebtById: Successfully ", resp)
	c.JSON(200, resp)
}

// @Router /debts	[get]
// @Summary GET DEBTS BY FILTER
// @Security  		BearerAuth
// @Description This method gets debts list by filter
// @Tags DEBTS
// @Accept json
// @Produce json
// @Param search query string false "Debt search"
// @Success 200 {object} models.GetDebtByFilterResp
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetDebtByFilter(c *gin.Context) {

	req := pb.GetDebtByFilterReq{}
	search := c.Query("search")

	unknown := strings.TrimSpace(search)
	if unknown == "" {
		req.Search = constanta
	}
	req.Search = unknown
	resp, err := h.service.GetDebtByFilter(context.Background(), &req)
	if err != nil {
		logger.Error("GetDebtByFilter: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("GetDebtByFilter: Successfully ", len(resp.Debt))
	c.JSON(200, resp)
}
