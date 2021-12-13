package controllers

import(
	"encoding/json"
	"io/ioutil"
	"kibby/product-cart/form"
	"kibby/product-cart/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductCartController struct{}

// AddProduct
func (pc ProductCartController) AddProductCart(c *gin.Context) {
	var req form.ProductCart
	var md models.ProductCartModel

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := json.Unmarshal(body, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	res, err := md.AddProductCart(req.UserID,
		req.Product,
		req.TotalPrice,
		)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"id": res})
	return
}