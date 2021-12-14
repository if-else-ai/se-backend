package form

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
    ID                  primitive.ObjectID      `json:"id" bson:"_id"`
    Status              string                  `json:"status" bson:"status"`
    Address             string                  `json:"address" bson:"address"`
    Detail              OrderDetail             `json:"detail" bson:"detail"`
    CustomerDetail      CustomerDetail          `json:"customerDetail" bson:"customerDetail"`
    TrackingNumber      string                  `json:"trackingNumber" bson:"trackingNumber"`
}

type OrderDetail struct {
    Product             []Product               `json:"products" bson:"products"`
    TotalPrice          float32                 `json:"totalPrice" bson:"totalPrice"`
    
}

type Product struct {
	ProductId		primitive.ObjectID			`json:"productId" bson:"productId"`
    Name         	string          			`json:"name" bson:"name"`
    Price        	float32         			`json:"price" bson:"price"`
    Quantity     	int32          				`json:"quantity" bson:"quantity"`
}

type CustomerDetail struct {
	UserId 			primitive.ObjectID 			`json:"userId" bson:"userId"`
    Name	        string    					`json:"name" bson:"name"`
    Tel     	  	string  					`json:"telNo" bson:"telNo"`
    Email 			string 						`json:"email" bson:"email"`
}

type OrderUpdate struct{
    Status              string                  `json:"status" bson:"status"`
    TrackingNumber      string                  `json:"trackingNumber" bson:"trackingNumber"`
}