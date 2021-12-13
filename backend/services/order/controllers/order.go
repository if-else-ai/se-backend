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

// //GetOrderByUserId
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

//GetOrderById
func (oc OrderController) GetOrderById(c *gin.Context) {
	var md models.OrderModel

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
		return
	}

	res, err := md.GetOrderByUserId(id)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

//CreateOrder
func (oc OrderController) CreateOrder(c *gin.Context){
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

	res, err := md.CreatOrder(req.Status,
		req.Address,
		req.Detail,
		req.CustomerDetail,
		req.TrackingNumber,)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"id": res})
	return
}