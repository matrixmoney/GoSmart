package controller

import (
    "smart-angel/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

type MarketScanController struct {
    service *service.MarketScanService
}

func NewMarketScanController(service *service.MarketScanService) *MarketScanController {
    return &MarketScanController{service: service}
}

func (ctrl *MarketScanController) FindOptionsByCompany(c *gin.Context) {
    // Implement the method to handle request and return options by company
}

func (ctrl *MarketScanController) MarketScan(c *gin.Context) {
    ctrl.service.MarketScan()
    c.JSON(http.StatusOK, gin.H{"message": "Market scan triggered"})
}
