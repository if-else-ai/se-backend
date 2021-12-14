package server

import (
	"kibby/order/controllers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	var orderController controllers.OrderController

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": "pong",
		})
	})
	router.GET("/orders", orderController.GetOrder)
	router.GET("/order/:userId", orderController.GetOrderByUserId)
	router.GET("/orderId/:id", orderController.GetOrderById)
	router.POST("/order",orderController.CreateOrder)

	return router
}

