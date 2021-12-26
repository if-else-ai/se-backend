package server

import (
	"kibby/admin/controllers"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	var adminController controllers.AdminController

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": "pong",
		})
	})

	router.POST("/login", adminController.Login)
	router.POST("/admin", adminController.AddAdmin)
	router.GET("/admin", adminController.GetAllAdmin)
	router.GET("/admin/:id", adminController.GetAdminByID)
	router.PUT("/admin/:id", adminController.UpdateAdmin)
	router.PUT("/admin/:id/password", adminController.UpdatePassword)
	router.DELETE("/admin/:id", adminController.DeleteAdmin)

	return router
}
