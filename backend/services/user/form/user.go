package form

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password"`
	TelNo       string             `json:"telNo" bson:"telNo"`
	Address     string             `json:"address" bson:"address"`
	DateOfBirth primitive.DateTime `json:"dateOfBirth" bson:"dateOfBirth"`
	Gender      string             `json:"gender" bson:"gender"`
}

type RegisterForm struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	PasswordSalt string             `json:"passwordSalt" bson:"passwordSalt"`
}

type UserUpdate struct {
	Name        string             `json:"name" bson:"name"`
	Email       string             `json:"email" bson:"email"`
	TelNo       string             `json:"telNo" bson:"telNo"`
	Address     string             `json:"address" bson:"address"`
	DateOfBirth primitive.DateTime `json:"dateOfBirth" bson:"dateOfBirth"`
	Gender      string             `json:"gender" bson:"gender"`
}
type PasswordUpdate struct {
	Password string `json:"password" bson:"password"`
}
