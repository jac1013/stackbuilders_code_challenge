package main

import "github.com/swaggo/gin-swagger"
import "github.com/swaggo/gin-swagger/swaggerFiles"

import (
	"github.com/gin-gonic/gin"
	"crypto_bot/exchange_pullers/binance/fetchers"
	"net/http"
	_ "crypto_bot/docs"
	_ "crypto_bot/business/structures"

	"stackbuilders_pizza/persistence"
	persistence2 "stackbuilders_pizza/business/persistence"
	"stackbuilders_pizza/business/useCases"
)

var memory persistence2.OrderPersistence

func main() {
	memory = persistence.NewOnMemoryPersistence()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/orders/create", orderCreateHandler)
	r.GET("/orders/:id", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/orders/", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}

func orderCreateHandler(c *gin.Context) {

	order, _ := useCases.NewOrderOperation(memory).Create()
	c.String(http.StatusOK, json)

}
