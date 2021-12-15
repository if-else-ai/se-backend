package form

import (
	"github.com/omise/omise-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	UserID         string             `json:"userId" bson:"userId"`
	Status         string             `json:"status" bson:"status"`
	Address        string             `json:"address" bson:"address"`
	Detail         OrderDetail        `json:"detail" bson:"detail"`
	UserDetail     UserDetail         `json:"userDetail" bson:"userDetail"`
	TrackingNumber string             `json:"trackingNumber" bson:"trackingNumber"`
}

type OrderDetail struct {
	Product    []Product    `json:"product" bson:"product"`
	TotalPrice float32      `json:"totalPrice" bson:"totalPrice"`
	Payment    omise.Charge `json:"payment" bson:"payment"`
}

type Product struct {
	ProductId primitive.ObjectID `json:"productId" bson:"productId"`
	Name      string             `json:"name" bson:"name"`
	Price     float32            `json:"price" bson:"price"`
	Quantity  int32              `json:"quantity" bson:"quantity"`
	Option    []ProductOption    `json:"option" bson:"option"`
	Image     []string           `json:"image" bson:"image"`
}

type ProductOption struct {
	Name       string  `json:"name" bson:"name"`
	Select     string  `json:"select" bson:"select"`
	PriceAdded float32 `json:"priceAdded" bson:"priceAdded"`
}

type UserDetail struct {
	Name  string `json:"name" bson:"name"`
	Tel   string `json:"telNo" bson:"telNo"`
	Email string `json:"email" bson:"email"`
}

type OrderUpdate struct {
	Status         string `json:"status" bson:"status"`
	TrackingNumber string `json:"trackingNumber" bson:"trackingNumber"`
}

type CreateOrderForm struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	UserID         primitive.ObjectID `json:"userId" bson:"userId"`
	Status         string             `json:"status" bson:"status"`
	Address        string             `json:"address" bson:"address"`
	Detail         OrderDetail        `json:"detail" bson:"detail"`
	UserDetail     UserDetail         `json:"userDetail" bson:"userDetail"`
	TrackingNumber string             `json:"trackingNumber" bson:"trackingNumber"`
}

type CreateOrderResponse struct {
	ID      string       `json:"id"`
	Payment omise.Charge `json:"payment"`
}

type UpdateOrderStatusFrom struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Status         string             `json:"status" bson:"status"`
	PaymentID      string             `json:"paymentId" bson:"paymentId"`
	TrackingNumber string             `json:"trackingNumber" bson:"trackingNumber"`
}
