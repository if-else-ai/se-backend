package controllers

import (
	"encoding/json"
	"io/ioutil"
	"kibby/product/form"
	"kibby/product/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

// GetProducts
func (pc ProductController) GetProducts(c *gin.Context) {
	var md models.ProductModel

	res, err := md.GetProducts()
	if err != nil {
		panic(err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// AddProduct
func (pc ProductController) AddProduct(c *gin.Context) {
	var req form.Product
	var md models.ProductModel

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := json.Unmarshal(body, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	res, err := md.AddProduct(req.Name,
		req.Category,
		req.Price,
		req.Description,
		req.Quantity,
		req.Option,
		req.Tag)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"id": res})
	return
}
