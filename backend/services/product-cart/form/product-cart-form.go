package form

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductCartForm struct {
	UserID  primitive.ObjectID `json:"userId" bson:"userId"`
	Product []Products         `json:"products" bson:product`
}

type Products struct {
	Name         string          `json:"name" bson:"name"`
	Price        float32         `json:"price" bson:"price"`
	Quantity     int32           `json:"quantity" bson:"quantity"`
	OptionDetail []OptionDetails `json:"optionDetail" bson:"optionDetail"`
}

type OptionDetails struct {
	Name     string  `json:"name" bson:"name"`
	Select   string  `json:"select" bson:"select"`
	PriceAdd float32 `json:"priceAdd" bson:"priceAdd"`
}
