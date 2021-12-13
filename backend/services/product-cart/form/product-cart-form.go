package form

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductCart struct {
	UserID  primitive.ObjectID `json:"userId" bson:"userId"`
	Product []Products `json:"products" bson:"products"`
	TotalPrice float32 `json:"totalPrice" bson:totalPrice`
}

type Products struct {
    Name         string          `json:"name" bson:"name"`
    Price        float32         `json:"price" bson:"price"`
    Quantity     int32           `json:"quantity" bson:"quantity"`
    OptionDetail []OptionDetails `json:"optionDetail" bson:"optionDetail"`
}

type OptionDetails struct {
	Name     string  `json:"optionname" bson:"optionname"`
	Select   string  `json:"optionselect" bson:"optionselect"`
	Price float32 `json:"optionpriceAdd" bson:"optionpriceAdd"`
}
