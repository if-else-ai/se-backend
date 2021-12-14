package models

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kibby/order/database"
	"kibby/order/form"
	"math"
	"net/http"
	"net/url"
	"strings"

	"github.com/omise/omise-go"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderModel struct{}

//GetOrder
func (o OrderModel) GetOrder() ([]form.Order, error) {
	coll, err := database.GetDB()
	if err != nil {
		return []form.Order{}, err
	}

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return []form.Order{}, errors.Wrap(err, "failed to find Order")
	}
	defer cursor.Close(context.TODO())

	var results []form.Order
	if err := cursor.All(context.TODO(), &results); err != nil {
		return []form.Order{}, errors.Wrap(err, "failed to get order")
	}
	return results, nil
}

//GetOrderByUserId
func (o OrderModel) GetOrderByUserId(userId primitive.ObjectID) ([]form.Order, error) {
	coll, err := database.GetDB()
	if err != nil {
		return []form.Order{}, err
	}

	cursor, err := coll.Find(context.TODO(), bson.M{"userId": userId})
	if err != nil {
		return []form.Order{}, errors.Wrap(err, "failed to find Order")
	}
	defer cursor.Close(context.TODO())

	var results []form.Order
	if err := cursor.All(context.TODO(), &results); err != nil {
		return []form.Order{}, errors.Wrap(err, "failed to get order")
	}
	return results, nil

}

//GetOrderById
func (o OrderModel) GetOrderById(id primitive.ObjectID) (form.Order, error) {
	coll, err := database.GetDB()
	if err != nil {
		return form.Order{}, err
	}

	var result form.Order
	if err := coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result); err != nil {
		return form.Order{}, errors.Wrap(err, "failed to get OrderById")
	}

	return result, nil
}

// CreateOrder
func (o OrderModel) CreatOrder(userId primitive.ObjectID,
	status string,
	address string,
	detail form.OrderDetail,
	userDetail form.UserDetail,
	trackingNumber string) (form.CreateOrderResponse, error) {
	coll, err := database.GetDB()
	if err != nil {
		return form.CreateOrderResponse{}, err
	}

	data := url.Values{
		"amount":            {fmt.Sprintf("%f", math.Round(float64(detail.TotalPrice*100)))},
		"currency":          {"THB"},
		"source[type]":      {"promptpay"},
		"metadata[user_id]": {userId.Hex()},
	}

	omiseReq, err := http.NewRequest("POST", "https://api.omise.co/charges", strings.NewReader(data.Encode()))
	if err != nil {
		return form.CreateOrderResponse{}, errors.Wrap(err, "failed to create omise request")
	}

	omiseReq.Header.Add("Authorization", "Basic c2tleV90ZXN0XzVwbmlxZGJmZXU1bmplNDVkYnY6OHhYTSZrcGJRQDlad0VsYmlpWlIzVkhP")

	omiseRes, err := http.DefaultClient.Do(omiseReq)
	if err != nil {
		return form.CreateOrderResponse{}, errors.Wrap(err, "failed to create omise request")
	}

	omiseResData, err := ioutil.ReadAll(omiseRes.Body)
	if err != nil {
		return form.CreateOrderResponse{}, errors.Wrap(err, "failed to read omise response")
	}

	var response form.CreateOrderResponse
	if err := json.Unmarshal(omiseResData, &response.Payment); err != nil {
		return form.CreateOrderResponse{}, errors.Wrap(err, "failed to unmarshal omise response")
	}

	// Document to insert
	doc := form.CreateOrderForm{
		ID:      primitive.NewObjectID(),
		UserID:  userId,
		Status:  status,
		Address: address,
		Detail: form.OrderDetail{
			TotalPrice: detail.TotalPrice,
			Product:    detail.Product,
			Payment:    response.Payment,
		},
		UserDetail:     userDetail,
		TrackingNumber: trackingNumber,
	}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return form.CreateOrderResponse{}, errors.Wrap(err, "failed to insert order")
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	response.ID = id

	return response, nil

}

// UpdateOrderStatus&Tracking
func (o OrderModel) UpdateOrderStatusAndTracking(id primitive.ObjectID,
	status string,
	paymentID string,
	trackingNumber string) (string, error) {
	coll, err := database.GetDB()
	if err != nil {
		return "", err
	}

	// Document to update
	// doc := form.OrderUpdate{
	// 	Status: status,
	// }

	if status == "Shipping" {
		if _, err := coll.UpdateByID(context.TODO(), id, bson.D{{"$set", bson.D{{"trackingNumber", trackingNumber}}}}); err != nil {
			return "", errors.Wrap(err, "failed to update status")
		}
	} else if status == "Paid" {
		req, err := http.NewRequest("POST", "https://api.omise.co/charges/"+paymentID+"/mark_as_paid", nil)
		if err != nil {
			return "", errors.Wrap(err, "failed to create omise request")
		}

		req.Header.Add("Authorization", "Basic c2tleV90ZXN0XzVwbmlxZGJmZXU1bmplNDVkYnY6OHhYTSZrcGJRQDlad0VsYmlpWlIzVkhP")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", errors.Wrap(err, "failed to create omise request")
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", errors.Wrap(err, "failed to read omise response")
		}

		var payment omise.Charge
		if err := json.Unmarshal(body, &payment); err != nil {
			return "", errors.Wrap(err, "failed to unmarshal omise response")
		}

		if _, err := coll.UpdateByID(context.TODO(), id, bson.D{{"$set", bson.D{{"detail.payment", payment}}}}); err != nil {
			return "", errors.Wrap(err, "failed to update status")
		}
		
	}

	// update := bson.D{
	// 	{"$set", bson.D{
	// 		{"status", status},

	// 	}}
	// }
	if _, err := coll.UpdateByID(context.TODO(), id, bson.D{{"$set", bson.D{{"status", status}}}}); err != nil {
		return "", errors.Wrap(err, "failed to update status")
	}

	return "update success", nil
}

//deleteOrder
func (o OrderModel) DeleteOrder(id primitive.ObjectID) (string, error) {
	coll, err := database.GetDB()
	if err != nil {
		return "", err
	}

	filter := bson.M{"_id": id}

	if _, err := coll.DeleteOne(context.TODO(), filter); err != nil {
		return "", errors.Wrap(err, "failed to delete document")
	}
	return "delete success", nil

}
