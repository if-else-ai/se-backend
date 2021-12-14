package models

import (
	"context"
	"crypto/rand"
	"io"
	"kibby/user/auth"
	"kibby/user/database"
	"kibby/user/form"
	"net/http"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct{}

// Register
func (u UserModel) Register(email string, password string) (form.RegisterResponseForm, int, error) {
	coll, err := database.GetDB()
	if err != nil {
		return form.RegisterResponseForm{}, http.StatusInternalServerError, err
	}

	// Check if email already exists
	findEmailRes := coll.FindOne(context.TODO(), bson.M{"email": email})
	if findEmailRes.Err() == nil {
		return form.RegisterResponseForm{}, http.StatusBadRequest, errors.New("email already exists")
	}

	// Generate password salt
	salt := make([]byte, 64)
	_, err = io.ReadFull(rand.Reader, salt)
	if err != nil {
		return form.RegisterResponseForm{}, http.StatusInternalServerError, errors.Wrap(err, "failed to generate password salt")
	}

	// Argon2 hash password
	saltedPassword := password + string(salt)
	hashedPassword, err := argon2id.CreateHash(saltedPassword, argon2id.DefaultParams)
	if err != nil {
		return form.RegisterResponseForm{}, http.StatusInternalServerError, errors.Wrap(err, "failed to hash password")
	}

	// Document to insert
	doc := form.RegisterForm{
		ID:           primitive.NewObjectID(),
		Email:        email,
		Password:     hashedPassword,
		PasswordSalt: string(salt),
	}

	res, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return form.RegisterResponseForm{}, http.StatusInternalServerError, errors.Wrap(err, "failed to insert document")
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()

	// Create token and auth
	tokenDetails, err := auth.CreateToken()
	if err != nil {
		return form.RegisterResponseForm{}, http.StatusInternalServerError, err
	}
	if err := auth.CreateAuth(id, tokenDetails); err != nil {
		return form.RegisterResponseForm{}, http.StatusInternalServerError, err
	}

	return form.RegisterResponseForm{
		ID:    id,
		Token: tokenDetails.AccessToken,
	}, http.StatusOK, nil
}

// AddUser
func (u UserModel) AddUser(name string,
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

	id := result.InsertedID.(primitive.ObjectID).Hex()

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

//GetUserByID
func (u UserModel) GetUserByID(id primitive.ObjectID) (form.User, error) {
	coll, err := database.GetDB()
	if err != nil {
		return form.User{}, err
	}

	var result form.User
	if err := coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result); err != nil {
		return form.User{}, errors.Wrap(err, "failed to get product")
	}
	return result, nil
}

//UpdateUser
func (u UserModel) UpdateUser(id primitive.ObjectID,
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
	doc := form.UserUpdate{
		Name:        name,
		Email:       email,
		TelNo:       telNo,
		Address:     address,
		DateOfBirth: primitive.NewDateTimeFromTime(dt),
		Gender:      gender,
	}
	update := bson.D{{"$set", doc}}
	if _, err := coll.UpdateByID(context.TODO(), id, update); err != nil {
		return "", errors.Wrap(err, "failed to update document")
	}

	return "update success", nil
}

//UpdatePassword
func (u UserModel) UpdatePassword(id primitive.ObjectID,
	ps string) (string, error) {

	coll, err := database.GetDB()
	if err != nil {
		return "", err
	}
	doc := form.PasswordUpdate{
		Password: ps,
	}
	update := bson.D{{"$set", doc}}
	if _, err := coll.UpdateByID(context.TODO(), id, update); err != nil {
		return "", errors.Wrap(err, "failed to update document")
	}
	return "update password success", nil
}

//delete
func (u UserModel) DeleteUser(id primitive.ObjectID) (string, error) {
	coll, err := database.GetDB()
	if err != nil {
		return "", err
	}

	filter := bson.M{"_id":id}

	if _, err := coll.DeleteOne(context.TODO(), filter); err != nil {
		return "", errors.Wrap(err, "failed to delete document")
	}
	return "delete success", nil
}
