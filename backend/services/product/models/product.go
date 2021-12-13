package models

import (
	"context"
	"kibby/product/database"
	"kibby/product/form"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductModel struct{}

// GetProducts
func (p ProductModel) GetProducts() ([]form.Product, error) {
	coll, err := database.GetDB()
	if err != nil {
		return []form.Product{}, err
	}

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return []form.Product{}, errors.Wrap(err, "failed to find products")
	}
	defer cursor.Close(context.TODO())

	var results []form.Product
	if err := cursor.All(context.TODO(), &results); err != nil {
		return []form.Product{}, errors.Wrap(err, "failed to get products")
	}

	return results, nil
}

// GetProductByID
func (p ProductModel) GetProductByID(id primitive.ObjectID) (form.Product, error) {
	coll, err := database.GetDB()
	if err != nil {
		return form.Product{}, err
	}

	var result form.Product
	if err := coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result); err != nil {
		return form.Product{}, errors.Wrap(err, "failed to get product")
	}

	return result, nil
}

// AddProduct
func (p ProductModel) AddProduct(name string,
	category string,
	price float32,
	description string,
	quantity int32,
	option []form.ProductOption,
	tag []string) (string, error) {
	coll, err := database.GetDB()
	if err != nil {
		return "", err
	}

	// Document to insert
	doc := form.Product{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Category:    category,
		Price:       price,
		Description: description,
		Quantity:    quantity,
		Option:      option,
		Tag:         tag,
	}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}
