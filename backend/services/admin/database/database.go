package database

import (
	"context"
	"fmt"
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

	fmt.Printf("connectionURI: %v\n", connectionURI)

	// Check the connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, errors.Wrap(err, "failed to ping mongo")
	}

	coll := client.Database("kibby").Collection("admin")

	return coll, err
}
