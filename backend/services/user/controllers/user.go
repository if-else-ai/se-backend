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

//UpdateUser
func (uc UserController) UpdateUser(c *gin.Context){
	var req form.User
	var md models.UserModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := md.UpdateUser(req.ID,
		req.Name,
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

//DeleteUser
func (us UserController) DeleteUser(c *gin.Context){
	var req form.User
	var md models.UserModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := md.DeleteUser(req.ID,)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": res})
	return 

}