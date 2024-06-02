// dbconnect/dbconnect.go

package dbconnect

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionString = "mongodb+srv://samiran4209:Samiran123@cluster0.ungeocm.mongodb.net/"
	dbName           = "Auth"
	colName          = "goAuth"
)

// ConnectMongoDB connects to MongoDB and returns the collection object
func ConnectMongoDB() *mongo.Collection {
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Create a new MongoDB client
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Create a context with a timeout to connect to the database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect the client
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get a handle for your collection
	collection := client.Database(dbName).Collection(colName)

	log.Println("Connected to MongoDB and initialized collection")

	return collection
}
