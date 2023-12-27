// config/database.go
package config

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBConnector struct represents a MongoDB connection
type DBConnector struct {
	client     *mongo.Client
	database   *mongo.Database
	ctx        context.Context
}


const DatabaseName = "my_app"

// NewDBConnector creates a new instance of DBConnector
func NewDBConnector() *DBConnector {
	return &DBConnector{
		ctx: context.TODO(),
	}
}

// Connect establishes a connection to MongoDB and sets the client field
func (db *DBConnector) Connect() error {
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI("mongodb+srv://dokhaihung:dokhaihung@cluster0.wzjck.mongodb.net/?retryWrites=true&w=majority")
	
	client, err := mongo.Connect(db.ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = client.Ping(db.ctx, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("Connected to MongoDB!")

	db.client = client
	db.database = client.Database(DatabaseName)

	return nil
}

// GetClient returns the MongoDB client
func (db *DBConnector) GetClient() *mongo.Client {
	return db.client
}


func (db *DBConnector) GetDatabase() *mongo.Database {
	return db.database
}