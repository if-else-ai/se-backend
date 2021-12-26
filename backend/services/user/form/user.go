package form

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	PasswordSalt string             `json:"passwordSalt" bson:"passwordSalt"`
	TelNo        string             `json:"telNo" bson:"telNo"`
	Address      string             `json:"address" bson:"address"`
	DateOfBirth  primitive.DateTime `json:"dateOfBirth" bson:"dateOfBirth"`
	Gender       string             `json:"gender" bson:"gender"`
}

type RegisterForm struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	PasswordSalt []byte             `json:"passwordSalt" bson:"passwordSalt"`
}

type RegisterResponseForm struct {
	ID string `json:"id"`
}

type LoginForm struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type LoginResultForm struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Password     string             `json:"password" bson:"password"`
	PasswordSalt []byte             `json:"passwordSalt" bson:"passwordSalt"`
}

type LoginResponseForm struct {
	ID string `json:"id"`
}

type UserUpdate struct {
	Name        string             `json:"name" bson:"name"`
	Email       string             `json:"email" bson:"email"`
	TelNo       string             `json:"telNo" bson:"telNo"`
	Address     string             `json:"address" bson:"address"`
	DateOfBirth primitive.DateTime `json:"dateOfBirth" bson:"dateOfBirth"`
	Gender      string             `json:"gender" bson:"gender"`
}

type PasswordUpdateForm struct {
	Password     string `json:"password" bson:"password"`
	PasswordSalt []byte `json:"passwordSalt" bson:"passwordSalt"`
}

type PasswordUpdateRequestForm struct {
	OldPassword string `json:"oldPassword" bson:"oldPassword"`
	NewPassword string `json:"newPassword" bson:"newPassword"`
}
