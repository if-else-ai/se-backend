package form

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductCart struct {
	UserID  primitive.ObjectID `json:"userId" bson:"userId"`
	Product []Product `json:"product" bson:"product"`
	TotalPrice float32 `json:"totalPrice" bson:totalPrice`
}

type Product struct {
    Name         string          `json:"name" bson:"name"`
    Price        float32         `json:"price" bson:"price"`
    Quantity     int32           `json:"quantity" bson:"quantity"`
    OptionDetail []OptionDetails `json:"optionDetail" bson:"optionDetail"`
}

type OptionDetails struct {
	Name     string  `json:"name" bson:"name"`
	Select   string  `json:"select" bson:"select"`
	Price 	 float32 `json:"priceAdded" bson:"priceAdded"`
}

