package models

import (
	"context"
	"fmt"
	"kibby/admin/database"
	"kibby/admin/form"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminModel struct{}

// AddAdmin
func (a AdminModel) AddAdmin(name string,
	email string,
	password string,
	telNo string,
	address string,
	dateOfBirth string,
	gender string) (string, error) {

	coll, err := database.GetDB()
	if err != nil {
		return "", err
	}

	dt, _ := time.Parse("2006-01-02", dateOfBirth)

	// Document
	doc := form.Admin{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Email:       email,
		Password:    password,
		TelNo:       telNo,
		Address:     address,
		DateOfBirth: primitive.NewDateTimeFromTime(dt),
		Gender:      gender,
	}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return "", errors.Wrap(err, "failed to insert document")
	}

	id := fmt.Sprint(result.InsertedID)

	return id, nil
}

// GetUsers
func (a AdminModel) GetAdmins() ([]form.Admin, error) {
	coll, err := database.GetDB()
	if err != nil {
		return []form.Admin{}, err
	}

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return []form.Admin{}, err
	}
	defer cursor.Close(context.TODO())

	var results []form.Admin
	if err = cursor.All(context.TODO(), &results); err != nil {
		return []form.Admin{}, err
	}
	return results, nil
}
//GetUserByID
func (a AdminModel) GetAdminByID(id primitive.ObjectID) (form.Admin, error) {
	coll, err := database.GetDB()
	if err != nil {
		return form.Admin{}, err
	}

	var result form.Admin
	if err := coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result); err != nil {
		return form.Admin{}, errors.Wrap(err, "failed to get product")
	}
	return result, nil
}

//UpdateUser
func (a AdminModel) UpdateAdmin(id primitive.ObjectID,
	name string,
	email string,
	telNo string,
	address string,
	dateOfBirth string,
	gender string) (string, error) {

	coll, err := database.GetDB()
	if err != nil {
		return "", err
	}
	dt, _ := time.Parse("2006-01-02", dateOfBirth)

	//Document
	doc := form.AdminUpdate{
		Name:        name,
		Email:		 email,
		TelNo:       telNo,
		Address:     address,
		DateOfBirth: primitive.NewDateTimeFromTime(dt),
		Gender:      gender,
	}
	update := bson.D{{"$set",doc}}
	if _ , err := coll.UpdateByID(context.TODO(),id,update); err != nil {
		return "", errors.Wrap(err, "failed to update document")
	}

	return "update success",nil
}

//UpdatePassword
func (a AdminModel) UpdatePassword(id primitive.ObjectID,
	ps string) (string,error){

	coll, err := database.GetDB()
	if err != nil {
		return "", err
	}
	doc:= form.PasswordUpdate{
		Password: ps,
	}
	update := bson.D{{"$set",doc}}
	if _ , err := coll.UpdateByID(context.TODO(),id,update); err != nil {
		return "", errors.Wrap(err, "failed to update document")
	}
	return "update password success",nil
}

//delete
func (a AdminModel) DeleteAdmin(id primitive.ObjectID) (string, error){
	coll, err := database.GetDB()
	if err != nil {
		return "", err
	}

	filter := bson.D{{"_id",id}}

	if _ ,err := coll.DeleteOne(context.TODO(),filter); err != nil{
		return "", errors.Wrap(err, "failed to delete document")
	}
	return "delete success",nil
}
