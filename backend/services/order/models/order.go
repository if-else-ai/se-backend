package models

import (
	"context"
	"kibby/order/database"
	"kibby/order/form"

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
func (o OrderModel) GetOrderByUserId(id primitive.ObjectID) (form.CustomerDetail, error) {
	coll, err := database.GetDB()
	if err != nil {
		return form.CustomerDetail{}, err
	}

	var result form.CustomerDetail
	if err := coll.FindOne(context.TODO(), bson.M{"userId": id}).Decode(&result); err != nil {
		return form.CustomerDetail{}, errors.Wrap(err, "failed to get OrderByUserId")
	}
	return result, nil
}

//CreateOrder
func (o OrderModel)	CreatOrder(status string,
	address string,
	detail form.OrderDetail,
	customDetail form.CustomerDetail,
	trackingNumber string,
	) (string, error) {
		coll,err := database.GetDB()
		if err != nil {
			return "", err
		}

		doc := form.Order{
			ID: primitive.NewObjectID(),
			Status: status,
			Address: address,
			Detail: detail,
			CustomerDetail:customDetail,
			TrackingNumber: trackingNumber,
		}
		result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
		
	}

//GetOrderById
func (o OrderModel) GetOrderById(id primitive.ObjectID) (form.Order, error){
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

//UpdateOrderStatus&Tracking
func (o OrderModel)  UpdateOrderStatusAndTracking(id primitive.ObjectID,
	status string,
	address string,
	detail form.OrderDetail,
	customDetail form.CustomerDetail,
	trackingNumber string,) (string, error){
	coll,err := database.GetDB()
	if err != nil {
		return "",err
	}
		doc := form.OrderUpdate{
				Status: status,}

		if status =="shipping"{
			doc = form.OrderUpdate{
			Status: status,
			TrackingNumber: trackingNumber,
			}
		}
		

	update := bson.D{{"$set",doc}}
	if _ , err := coll.UpdateByID(context.TODO(),id,update); err != nil {
		return "", errors.Wrap(err, "failed to update status")
	}

	return "update success",nil

}