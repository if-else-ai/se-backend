package models

import (
	"context"
	"kibby/user/database"
	"kibby/user/form"

	"go.mongodb.org/mongo-driver/bson"
)

type UserModel struct{}

// AddUser
func (u UserModel) AddUser(name string,
	email string,
	password string,
	gender string) (interface{}, error) {

	coll, err := database.GetDB()
	if err != nil {
		return "", err
	}

	doc := form.User{
		Name:     name,
		Email:    email,
		Password: password,
		Gender:   gender,
	}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return "", err
	}

	id := result.InsertedID

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
