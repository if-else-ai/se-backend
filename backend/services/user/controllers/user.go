package controllers

import (
	"kibby/user/form"
	"kibby/user/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct{}

// Register
func (uc UserController) Register(c *gin.Context) {
	var req form.RegisterForm
	var md models.UserModel

	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, statusCode, err := md.Register(req.Email, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Login
func (uc UserController) Login(c *gin.Context) {
	var req form.LoginForm
	var md models.UserModel

	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, statusCode, err := md.Login(req.Email, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(statusCode, res)
}

// AddUser
func (uc UserController) AddUser(c *gin.Context) {
	var req form.User
	var md models.UserModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := md.AddUser(req.Name,
		req.Email,
		req.Password,
		req.TelNo,
		req.Address,
		req.DateOfBirth.Time().Format("2006-01-02"),
		req.Gender)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": res})
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

//GetUsersByID
func (uc UserController) GetUsersByID(c *gin.Context) {
	var md models.UserModel

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
		return
	}

	res, err := md.GetUserByID(id)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

//UpdateUser
func (uc UserController) UpdateUser(c *gin.Context) {
	var req form.User
	var md models.UserModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
		return
	}

	res, err := md.UpdateUser(id,
		req.Name,
		req.Email,
		req.TelNo,
		req.Address,
		req.DateOfBirth.Time().Format("2006-01-02"),
		req.Gender)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": res})
	return

}

// UpdatePassword
func (us UserController) UpdatePassword(c *gin.Context) {
	var req form.PasswordUpdateRequestForm
	var md models.UserModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
		return
	}

	res, statusCode, err := md.UpdatePassword(id, req.OldPassword, req.NewPassword)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(statusCode, gin.H{"message": res})
	return
}

//DeleteUser
func (us UserController) DeleteUser(c *gin.Context) {

	var md models.UserModel

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
		return
	}

	res, err := md.DeleteUser(id)
	if err != nil {
		panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": res})
	return

}
