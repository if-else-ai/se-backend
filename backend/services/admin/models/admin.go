package models

import (
	"kibby/admin/database"
	"kibby/admin/form"

	"context"
	"crypto/rand"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AdminModel struct{}

// Register
func (a AdminModel) Register(email string, password string) (form.RegisterResponseForm, int, error) {
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

	// Hash password
	saltedPassword := password + string(salt)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), 10)
	if err != nil {
		return form.RegisterResponseForm{}, http.StatusInternalServerError, errors.Wrap(err, "failed to hash password")
	}

	// Document to insert
	doc := form.RegisterForm{
		ID:           primitive.NewObjectID(),
		Email:        email,
		Password:     string(hashedPassword),
		PasswordSalt: salt,
	}

	res, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return form.RegisterResponseForm{}, http.StatusInternalServerError, errors.Wrap(err, "failed to insert document")
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()

	return form.RegisterResponseForm{
		ID: id,
	}, http.StatusOK, nil
}

// Login
func (a AdminModel) Login(email string, password string) (form.LoginResponseForm, int, error) {
	coll, err := database.GetDB()
	if err != nil {
		return form.LoginResponseForm{}, http.StatusInternalServerError, err
	}

	FindEmailRes := coll.FindOne(context.TODO(), bson.M{"email": email})
	if FindEmailRes.Err() != nil {
		return form.LoginResponseForm{}, http.StatusBadRequest, errors.New("email not found")
	}

	var user form.LoginResultForm
	if err := FindEmailRes.Decode(&user); err != nil {
		return form.LoginResponseForm{}, http.StatusInternalServerError, errors.Wrap(err, "failed to decode document")
	}

	// Check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+string(user.PasswordSalt))); err != nil {
		return form.LoginResponseForm{}, http.StatusUnauthorized, errors.New("password is incorrect")
	}

	id := user.ID.Hex()

	return form.LoginResponseForm{
		ID: id,
	}, http.StatusOK, nil
}

// AddAdmin
func (a AdminModel) AddAdmin(email string, password string) (form.RegisterResponseForm, int, error) {
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

	// Hash password
	saltedPassword := password + string(salt)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), 10)
	if err != nil {
		return form.RegisterResponseForm{}, http.StatusInternalServerError, errors.Wrap(err, "failed to hash password")
	}

	// Document to insert
	doc := form.RegisterForm{
		ID:           primitive.NewObjectID(),
		Email:        email,
		Password:     string(hashedPassword),
		PasswordSalt: salt,
	}

	res, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return form.RegisterResponseForm{}, http.StatusInternalServerError, errors.Wrap(err, "failed to insert document")
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()

	return form.RegisterResponseForm{
		ID: id,
	}, http.StatusOK, nil
}

// GetAllAdmin
func (a AdminModel) GetAllAdmin() ([]form.Admin, error) {
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

//GetAdminByID
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

//UpdateAdmin
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

// UpdatePassword
func (a AdminModel) UpdatePassword(id primitive.ObjectID,
	oldPassword string, newPassword string) (string, int, error) {

	coll, err := database.GetDB()
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	var passwordUpdateForm form.PasswordUpdateForm
	if err := coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&passwordUpdateForm); err != nil {
		return "", http.StatusInternalServerError, errors.Wrap(err, "failed to get password update form")
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(passwordUpdateForm.Password),
		[]byte(oldPassword+string(passwordUpdateForm.PasswordSalt))); err != nil {
		return "", http.StatusUnauthorized, errors.Wrap(err, "wrong password")
	}

	// Generate password salt
	salt := make([]byte, 64)
	_, err = io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", http.StatusInternalServerError, errors.Wrap(err, "failed to generate password salt")
	}

	// Hash password
	saltedPassword := newPassword + string(salt)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), 10)
	if err != nil {
		return "", http.StatusInternalServerError, errors.Wrap(err, "failed to hash password")
	}

	doc := form.PasswordUpdateForm{
		Password:     string(hashedPassword),
		PasswordSalt: salt,
	}
	update := bson.D{{"$set", doc}}

	if _, err := coll.UpdateByID(context.TODO(), id, update); err != nil {
		return "", http.StatusInternalServerError, errors.Wrap(err, "failed to update document")
	}

	return "update password success", http.StatusOK, nil
}

//delete
func (a AdminModel) DeleteAdmin(id primitive.ObjectID) (string, error) {
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
