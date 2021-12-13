package models

import (
	"context"
	"kibby/product-cart/form"
	"kibby/product-cart/database"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductCartModel struct {
}

//AddProductCart
func (p ProductCartModel) AddProductCart(id primitive.ObjectID,
	product []form.Product,
	totalPrice float32) (string, error) {

	coll, err := database.GetDB()
	if err != nil {
		return "", errors.Wrap(err, "failed to GetDB")
	}
	
	doc := form.ProductCart{
		UserID: id,
		Product: product,
		TotalPrice: totalPrice,
	}

	_ , err = coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return "", errors.Wrap(err, "failed to insert document")
	}
	return "Add Product-cart success ", nil

}
