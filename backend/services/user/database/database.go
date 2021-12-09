package database

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionURI string

// Init
func Init() {
	connectionURI = os.Getenv("MONGO_URI")
}

// GetDB
func GetDB() (*mongo.Collection, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionURI))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to mongo")
	}

	coll := client.Database("kibby").Collection("user")

	return coll, err
}
