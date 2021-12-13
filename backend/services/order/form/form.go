package form

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
    ID                  primitive.ObjectID      `json:"id" bson:"_id"`
    Status              string                  `json:"name" bson:"name"`
    Address             string                  `json:"category" bson:"category"`
    Detail              []OrderDetail           `json:"price" bson:"price"`
    CustomerDetail      []CustomerDetail        `json:"description" bson:"description"`
    TrackingNumber      string                  `json:"tag" bson:"tag"`
}

type OrderDetail struct {
    Product             []Product               `json:"product" bson:"product"`
    TotalPrice          float32                 `json:"totalPrice" bson:"totaalPrice"`
    
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
    Tel     	  	string  					`json:"list" bson:"list"`
    Email 			string 						`json:"priceAdded" bson:"priceAdded"`
}
