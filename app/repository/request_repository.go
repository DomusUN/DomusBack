package repository

import (
	"DomusBack/app/domain"
	"DomusBack/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type RequestMetadata interface {
	Create(uvm *domain.RequestMetadata) (primitive.ObjectID, error)
	ChangeState(primitive.ObjectID, int) (primitive.ObjectID, error)
}

type RequestMetadataMongo struct{}

var collectionRequest = database.OpenCollection(database.Client, "requests")

func (u RequestMetadataMongo) Create(uvm *domain.RequestMetadata) (primitive.ObjectID, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	// Check the id of the client
	var user *domain.UsersMetadata
	err := collectionUser.FindOne(ctx, bson.M{"_id": uvm.IdClient}).Decode(&user)
	if err != nil || user.Role != 0 {
		log.Printf("Client user not found: %v", err)
		return primitive.NilObjectID, err
	}

	//Check the id of the worker
	err = collectionUser.FindOne(ctx, bson.M{"_id": uvm.IdWorker}).Decode(&user)
	if err != nil || user.Role != 1 {
		log.Printf("Worker user not found: %v", err)
		return primitive.NilObjectID, err
	}

	// If both exists add the ID
	uvm.ID = primitive.NewObjectID()

	result, err := collectionRequest.InsertOne(ctx, uvm)
	if err != nil {
		log.Printf("Could not create userdata: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, err
}

func (u RequestMetadataMongo) ChangeState(idRequest primitive.ObjectID, stateRequest int) (primitive.ObjectID, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	err := collectionRequest.FindOneAndUpdate(ctx, bson.M{"_id": idRequest}, bson.M{"$set": bson.M{"state": stateRequest}}, options.FindOneAndUpdate())

	if err.Err() != nil {
		log.Printf("Could not change the state: %v", err)
		return idRequest, err.Err()
	}

	return idRequest, err.Err()
}
