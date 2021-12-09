package controllers

import (
	"kibby/user/form"
	"kibby/user/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// AddUser
func (uc UserController) AddUser(c *gin.Context) {
	var req form.User
	var md models.UserModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := md.AddUser(req.Name, req.Email, req.Password, req.Gender)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// GetUsers
func (uc UserController) GetUsers(c *gin.Context) {
	var md models.UserModel

	res, err := md.GetUsers()
	if err != nil {
		panic(err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
