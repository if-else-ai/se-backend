package controllers

// import (
// 	"kibby/admin/form"
// 	"kibby/admin/models"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type AdminController struct{}

// // AddUser
// func (ac AdminController) AddAdmin(c *gin.Context) {
// 	var req form.Admin
// 	var md models.AdminModel

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	res, err := md.AddAdmin(req.Name,
// 		req.Email,
// 		req.Password,
// 		req.TelNo,
// 		req.Address,
// 		req.DateOfBirth.Time().Format("2006-01-02"),
// 		req.Gender)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"id": res})
// 	return
// }

// // GetAdmin
// func (ac AdminController) GetAdmins(c *gin.Context) {
// 	var md models.AdminModel

// 	res, err := md.GetAdmins()
// 	if err != nil {
// 		panic(err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// 	return
// }

// //GetUsersByID
// func (a AdminController) GetAdminByID(c *gin.Context) {
// 	var md models.AdminModel

// 	id, err := primitive.ObjectIDFromHex(c.Param("id"))
// 	if err != nil {
// 		panic(err)
// 		return
// 	}

// 	res, err := md.GetAdminByID(id)
// 	if err != nil {
// 		panic(err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// 	return
// }

// //UpdateAdmin
// func (ac AdminController) UpdateAdmin(c *gin.Context){
// 	var req form.Admin
// 	var md models.AdminModel

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
// 	res, err := md.UpdateAdmin(req.ID,
// 		req.Name,
// 		req.Email,
// 		req.TelNo,
// 		req.Address,
// 		req.DateOfBirth.Time().Format("2006-01-02"),
// 		req.Gender)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": res})
// 	return

// }

// //UpdatePassword
// func (ac AdminController) UpdatePassword(c *gin.Context){
// 	var req form.Admin
// 	var md models.AdminModel

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
// 	res, err := md.UpdatePassword(req.ID,req.Password)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": res})
// 	return

// }

// //DeleteUser
// func (ac AdminController) DeleteAdmin(c *gin.Context){
// 	var req form.Admin
// 	var md models.AdminModel

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
// 	res, err := md.DeleteAdmin(req.ID,)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": res})
// 	return

// }
