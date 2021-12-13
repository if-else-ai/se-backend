package models

import (
	"context"
	"fmt"
	"kibby/product-cart/form"
	"kibby/user/database"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductCartModel struct {
}

//AddProductCart
func (p ProductCartModel) AddProductCart(
	id primitive.ObjectID,
	product []form.ProductCart,
	totalPrice float32
	) (string, error) {

	coll, err := database.GetDB()
	if err != nil {
		return "", err
	}
	

	doc := form.ProductCart{
		UserID: id,
		product: product,
		TotalPrice: totalPrice,
	}

}
