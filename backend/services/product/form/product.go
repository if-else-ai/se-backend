package form

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Category    string             `json:"category" bson:"category"`
	Price       float32            `json:"price" bson:"price"`
	Description string             `json:"description" bson:"description"`
	Quantity    int32              `json:"quantity" bson:"quantity"`
	Option      []ProductOption    `json:"option" bson:"option"`
	Image       []string           `json:"image" bson:"image"`
	Tag         []string           `json:"tag" bson:"tag"`
}

type ProductOption struct {
	Name       string    `json:"name" bson:"name"`
	List       []string  `json:"list" bson:"list"`
	PriceAdded []float32 `json:"priceAdded" bson:"priceAdded"`
}

type AddProductForm struct {
	Name        string          `json:"name" bson:"name"`
	Category    string          `json:"category" bson:"category"`
	Price       float32         `json:"price" bson:"price"`
	Description string          `json:"description" bson:"description"`
	Quantity    int32           `json:"quantity" bson:"quantity"`
	Option      []ProductOption `json:"option" bson:"option"`
	Tag         []string        `json:"tag" bson:"tag"`
}

type UpdateProductForm struct {
	Name        string          `json:"name" bson:"name"`
	Category    string          `json:"category" bson:"category"`
	Price       float32         `json:"price" bson:"price"`
	Description string          `json:"description" bson:"description"`
	Quantity    int32           `json:"quantity" bson:"quantity"`
	Option      []ProductOption `json:"option" bson:"option"`
	Tag         []string        `json:"tag" bson:"tag"`
}
