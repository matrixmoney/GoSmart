package main

import (
    "smart-angel/controller"
    "smart-angel/service"
    "smart-angel/config"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/jasonlvhit/gocron"
)

func main() {
    router := gin.Default()
    
    config.LoadEnv()
    
    marketScanService := service.NewMarketScanService()
    controller := controller.NewMarketScanController(marketScanService)
    
    router.GET("/market/options", controller.FindOptionsByCompany)
    router.GET("/market/scan", controller.MarketScan)
    
    go func() {
        gocron.Every(5).Minutes().Do(marketScanService.MarketScan)
        <-gocron.Start()
    }()
    
    if err := router.Run(); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
