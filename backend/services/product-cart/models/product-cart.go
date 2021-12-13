package models

import (
	"kibby/product-cart/database"
	"kibby/product-cart/form"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductCartForm struct{}

//AddProductToCart
func (p ProductCartForm) AddProduct(
	userID primitive.ObjectID,
	productName string,
	productPrice float32,
	productQuantity int,
	nameDetail string,
	detailSelect string,
	detailPrice float32) (string, error) {

	coll, err := database.GetDB()
	if err != nil {
		return "", err
	}

	doc


	doc := form.OptionDetails{
		UserID:	userID,
		Product: 
		
	}

}
