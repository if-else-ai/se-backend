package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	// var adminController controllers.AdminController

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": "pong",
		})
	})

	// router.POST("/admin", adminController.AddAdmin)
	// router.GET("/admins", adminController.GetAdmins)
	// router.GET("/admin/:id", adminController.GetAdminByID)
	// router.PUT("/admin", adminController.UpdateAdmin)
	// router.PUT("/adminP", adminController.UpdatePassword)
	// router.DELETE("/admin", adminController.DeleteAdmin)

	return router
}
