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
	router.GET("/orderByUser/:userId", orderController.GetOrderByUserId)
	router.GET("/orderById/:id", orderController.GetOrderById)
	router.POST("/order", orderController.CreateOrder)
	router.PUT("/order", orderController.UpdateOrderStatusAndTracking)
	router.DELETE("/order/:id", orderController.DeleteOrderByID)
	router.DELETE("/order", orderController.DeleteAllOrders)

	return router
}
