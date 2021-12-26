package controllers

import (
	"kibby/admin/form"
	"kibby/admin/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminController struct{}

// Register
func (ac AdminController) Register(c *gin.Context) {
	var req form.RegisterForm
	var md models.AdminModel

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
func (ac AdminController) Login(c *gin.Context) {
	var req form.LoginForm
	var md models.AdminModel

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

// AddAdmin
func (ac AdminController) AddAdmin(c *gin.Context) {
	var req form.RegisterRequestForm
	var md models.AdminModel

	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, statusCode, err := md.AddAdmin(req.Email, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// GetAllAdmin
func (ac AdminController) GetAllAdmin(c *gin.Context) {
	var md models.AdminModel

	res, err := md.GetAllAdmin()
	if err != nil {
		panic(err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// GetAdminByID
func (ac AdminController) GetAdminByID(c *gin.Context) {
	var md models.AdminModel

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
		return
	}

	res, err := md.GetAdminByID(id)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// UpdateAdmin
func (ac AdminController) UpdateAdmin(c *gin.Context) {
	var req form.Admin
	var md models.AdminModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
		return
	}

	res, err := md.UpdateAdmin(id,
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
func (ac AdminController) UpdatePassword(c *gin.Context) {
	var req form.PasswordUpdateRequestForm
	var md models.AdminModel

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

// DeleteAdmin
func (ac AdminController) DeleteAdmin(c *gin.Context) {

	var md models.AdminModel

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
		return
	}

	res, err := md.DeleteAdmin(id)
	if err != nil {
		panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": res})
	return
}
