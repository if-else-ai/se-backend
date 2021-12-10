package models

import (
	"context"
	"fmt"
	"kibby/user/database"
	"kibby/user/form"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct{}

// AddUser
func (u UserModel) AddUser(name string,
	email string,
	password string,
	telNo string,
	address string,
	dateOfBirth string,
	gender string) (interface{}, error) {

	coll, err := database.GetDB()
	if err != nil {
		return "", err
	}

	dt, _ := time.Parse("2006-01-02", dateOfBirth)

	// Document
	doc := form.User{
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

	update := bson.D{{"$set", bson.D{
		{"name", name},
		{"telNo", telNo},
	}}}

	coll.UpdateByID(context.TODO(), id, update)

	return id, nil
}

// GetUsers
func (u UserModel) GetUsers() ([]form.User, error) {
	coll, err := database.GetDB()
	if err != nil {
		return []form.User{}, err
	}

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return []form.User{}, err
	}
	defer cursor.Close(context.TODO())

	var results []form.User
	if err = cursor.All(context.TODO(), &results); err != nil {
		return []form.User{}, err
	}

	return results, nil
}
