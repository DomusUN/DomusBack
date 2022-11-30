package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 20
	connectionStringTemplate = "mongodb+srv://%s:%s@%s"
)

var mc *mongo.Client
var database *mongo.Database

// DBinstance func
func DBinstance() *mongo.Database {
	/*
		errEnv := godotenv.Load(".env")

		if errEnv != nil {
			log.Fatal("Error loading .env file")
		}
	*/

	//username := os.Getenv("REPRODUCTION_DB_USER")
	//password := os.Getenv("REPRODUCTION_DB_PASSWORD")
	//clusterEndpoint := os.Getenv("REPRODUCTION_DB_ENDPOINT")

	connectionURI := fmt.Sprintf(connectionStringTemplate, "domus_admin", "Domus_pass_001", "clusterdomus.47hyudk.mongodb.net/?retryWrites=true&w=majority")

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	defer cancel()

	var err error
	mc, err = mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
	}

	err = mc.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	database = mc.Database("DomusDB")
	//collection1 = mc.Database("ClusterDomus").Collection("users")

	fmt.Println("Connected to MongoDB!")

	return database
}

// Client Database instance
var Client = DBinstance()

// OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Database, collectionName string) *mongo.Collection {

	var collection = client.Collection(collectionName)

	fmt.Println("Connected to collection!")

	return collection
}
