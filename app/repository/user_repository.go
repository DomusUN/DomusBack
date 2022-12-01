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

type UserMetadata interface {
	Create(metadata *domain.UsersMetadata) (primitive.ObjectID, error)
	GetUserById(primitive.ObjectID) (*domain.UsersMetadata, error)
	AddRoleClient(primitive.ObjectID, *domain.ClientMetadata) (primitive.ObjectID, error)
	AddRoleWorker(primitive.ObjectID, *domain.WorkerMetadata) (primitive.ObjectID, error)
}

type UserMetadataMongo struct{}

var collectionUser = database.OpenCollection(database.Client, "users")

func (u UserMetadataMongo) Create(uvm *domain.UsersMetadata) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	uvm.ID = primitive.NewObjectID()

	result, err := collectionUser.InsertOne(ctx, uvm)
	if err != nil {
		log.Printf("Could not create userdata: %v", err)
		return primitive.NilObjectID, err
	}

	oid := result.InsertedID.(primitive.ObjectID)

	return oid, err
}

func (u UserMetadataMongo) GetUserById(id primitive.ObjectID) (*domain.UsersMetadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()
	var user *domain.UsersMetadata
	err := collectionUser.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	return user, err
}

func (u UserMetadataMongo) AddRoleClient(id primitive.ObjectID, client *domain.ClientMetadata) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	//Delete previous role if exists
	err := collectionUser.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$unset": bson.M{"worker": ""}}, options.FindOneAndUpdate())

	if err.Err() != nil {
		log.Printf("Could not delete old role: %v", err)
		return id, err.Err()
	}

	err = collectionUser.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"client": client, "role": 0}}, options.FindOneAndUpdate())

	return id, err.Err()
}

func (u UserMetadataMongo) AddRoleWorker(id primitive.ObjectID, worker *domain.WorkerMetadata) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	//Delete previous role if exists
	err := collectionUser.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$unset": bson.M{"client": ""}}, options.FindOneAndUpdate())

	if err.Err() != nil {
		log.Printf("Could not delete old role: %v", err)
		return id, err.Err()
	}

	err = collectionUser.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"worker": worker, "role": 1}}, options.FindOneAndUpdate())

	return id, err.Err()
}
