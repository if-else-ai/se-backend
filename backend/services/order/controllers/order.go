package controllers

import (
	"encoding/json"
	"io/ioutil"
	"kibby/order/form"
	"kibby/order/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderController struct{}

//GetAllOrder
func (oc OrderController) GetOrder(c *gin.Context) {
	var md models.OrderModel
	res, err := md.GetOrder()
	if err != nil {
		panic(err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

//GetOrderByUserId
func (oc OrderController) GetOrderByUserId(c *gin.Context) {
	var md models.OrderModel

	userId, err := primitive.ObjectIDFromHex(c.Param("userId"))
	if err != nil {
		panic(err)
		return
	}

	res, err := md.GetOrderByUserId(userId)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// GetOrderById
func (oc OrderController) GetOrderById(c *gin.Context) {
	var md models.OrderModel
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
		return
	}

	res, err := md.GetOrderById(id)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// CreateOrder
func (oc OrderController) CreateOrder(c *gin.Context) {
	var req form.Order
	var md models.OrderModel

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
		return
	}
	if err := json.Unmarshal(body, &req); err != nil {
		panic(err)
		return
	}

	userId, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := md.CreatOrder(userId,
		req.Status,
		req.Address,
		req.Detail,
		req.UserDetail,
		req.TrackingNumber,
		req.ShipStatus)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
	return
}

// UpdateOrderStatusAndTracking
func (oc OrderController) UpdateOrderStatusAndTracking(c *gin.Context) {
	var req form.UpdateOrderStatusFrom
	var md models.OrderModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := md.UpdateOrderStatusAndTracking(req.ID,
		req.Status,
		req.PaymentID,
		req.ShipStatus,
		req.TrackingNumber,
		
	)
	if err != nil {
		panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": res})
	return

}

//Delete
func (oc OrderController) DeleteOrder(c *gin.Context) {

	var md models.OrderModel

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
		return
	}

	res, err := md.DeleteOrder(id)
	if err != nil {
		panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": res})
	return

}
