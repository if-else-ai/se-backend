package form

import "go.mongodb.org/mongo-driver/bson/primitive"

<<<<<<< HEAD
type ProductCart struct {
	UserID  primitive.ObjectID `json:"userId" bson:"userId"`
	Product []Products `json:"products" bson:"products"`
	TotalPrice float32 `json:"totalPrice" bson:totalPrice`
=======
type ProductCartForm struct {
	UserID     primitive.ObjectID `json:"userId" bson:"userId"`
	Product    []Products         `json:"products" bson:product`
	TotalPrice float32            `json:"totalPrice" bson:totalPrice`
>>>>>>> f0aa5b449eeab7636194a75f91d13bdd27bbb996
}

type Products struct {
    Name         string          `json:"name" bson:"name"`
    Price        float32         `json:"price" bson:"price"`
    Quantity     int32           `json:"quantity" bson:"quantity"`
    OptionDetail []OptionDetails `json:"optionDetail" bson:"optionDetail"`
}

type OptionDetails struct {
<<<<<<< HEAD
	Name     string  `json:"optionname" bson:"optionname"`
	Select   string  `json:"optionselect" bson:"optionselect"`
	Price float32 `json:"optionpriceAdd" bson:"optionpriceAdd"`
=======
	Name     string  `json:"name" bson:"name"`
	Select   string  `json:"select" bson:"select"` //Select Attibute
	PriceAdd float32 `json:"priceAdd" bson:"priceAdd"`
>>>>>>> f0aa5b449eeab7636194a75f91d13bdd27bbb996
}

