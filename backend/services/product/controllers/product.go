package controllers

import (
	"fmt"
	"kibby/product/form"
	"kibby/product/models"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct{}

var fileDst = "/opt/files/"

// GetProducts
func (pc ProductController) GetProducts(c *gin.Context) {
	var md models.ProductModel

	res, err := md.GetProducts()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// GetProductByID
func (pc ProductController) GetProductByID(c *gin.Context) {
	var md models.ProductModel

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	res, err := md.GetProductByID(id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// GetProductImage
func (pc ProductController) GetProductImage(c *gin.Context) {
	name := c.Param("name")

	data, err := os.ReadFile(fileDst + name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.Data(http.StatusOK, "", data)
}

// AddProduct
func (pc ProductController) AddProduct(c *gin.Context) {
	var req form.AddProductForm
	var md models.ProductModel

	if err := c.Bind(&req); err != nil {
		panic(err)
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
}

// UpdateProduct
func (pc ProductController) UpdateProduct(c *gin.Context) {
	var req form.UpdateProductForm
	var md models.ProductModel

	if err := c.Bind(&req); err != nil {
		panic(err)
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	if err = md.UpdateProduct(id,
		req.Name,
		req.Category,
		req.Price,
		req.Description,
		req.Quantity,
		req.Option,
		req.Tag); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// AddProductImage
func (pc ProductController) AddProductImage(c *gin.Context) {
	var req struct {
		ProductID string                  `form:"productId"`
		Image     []*multipart.FileHeader `form:"image"`
	}
	var md models.ProductModel

	if err := c.Bind(&req); err != nil {
		panic(err)
	}

	id, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		panic(err)
	}

	product, err := md.GetProductByID(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("product: %v\n", product)

	var imageRequest []string
	multipartForm, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}
	images := multipartForm.File["image"]

	i := len(product.Image) + 1
	fmt.Printf("images: %v %v\n", len(images), i)
	for _, image := range images {
		filename := strconv.Itoa(i) + "-" + strconv.Itoa(int(time.Now().Unix())) + filepath.Ext(image.Filename)
		fmt.Printf("filename: %v\n", filename)
		imageRequest = append(imageRequest, "http://139.59.219.217:5102/image/"+filename)
		if err := c.SaveUploadedFile(image, fileDst+filename); err != nil {
			panic(err)
		}
		i++
	}

	if err := md.AddProductImage(id, imageRequest); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"image": imageRequest})
}
