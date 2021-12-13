package server

import (
	"net/http"
	"kibby/product-cart/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	var productCartController controllers.ProductCartController

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.POST("/productCart", productCartController.AddProductCart)
	router.GET("/productCart/:id", productCartController.GetProductCartByID)
	

	return router
}
