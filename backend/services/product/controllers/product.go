package controllers

import (
	"encoding/json"
	"io/ioutil"
	"kibby/product/form"
	"kibby/product/models"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetProductByID
func (pc ProductController) GetProductByID(c *gin.Context) {
	var md models.ProductModel

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
		return
	}

	res, err := md.GetProductByID(id)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// GetProductImage
func (pc ProductController) GetProductImage(c *gin.Context) {
	name := c.Param("name")

	data, err := os.ReadFile("/opt/files/" + name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.Data(http.StatusOK, "", data)
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

// UploadProductImage
func (pc ProductController) UploadProductImage(c *gin.Context) {
	var req struct {
		Image []*multipart.FileHeader `form:"image"`
	}

	if err := c.ShouldBind(&req); err != nil {
		panic(err)
		return
	}

	var res []string

	multipartForm, err := c.MultipartForm()
	if err != nil {
		panic(err)
		return
	}
	for s, _ := range multipartForm.File {
		if s == "image" {
			for _, file := range req.Image {
				dst := "/opt/files/" + strconv.Itoa(int(time.Now().Unix())) + ".png"
				res = append(res, dst)
				if err := c.SaveUploadedFile(file, dst); err != nil {
					panic(err)
					return
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"image": res})
	return
}
