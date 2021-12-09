package server

import (
	"kibby/user/controllers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	var userController controllers.UserController

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": "pong",
		})
	})

	router.POST("/user", userController.AddUser)
	router.GET("/users", userController.GetUsers)

	return router
}
