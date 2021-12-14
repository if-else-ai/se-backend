package server

import (
	"kibby/product/controllers"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	var productController controllers.ProductController

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": "pong",
		})
	})
	router.GET("/products", productController.GetProducts)
	router.GET("/product/:id", productController.GetProductByID)
	router.POST("/product", productController.AddProduct)
	router.PUT("/product/:id", productController.UpdateProduct)
	router.GET("/image/:name", productController.GetProductImage)
	router.POST("/image", productController.AddProductImage)

	return router
}
