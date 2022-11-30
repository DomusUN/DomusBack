package repository

import (
	"DomusBack/app/domain"
	"DomusBack/database"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type UserMetadata interface {
	Create(metadata *domain.UsersMetadata) (primitive.ObjectID, error)
}

type UserMetadataMongo struct{}

var collection = database.OpenCollection(database.Client, "users")

func (u UserMetadataMongo) Create(uvm *domain.UsersMetadata) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	uvm.ID = primitive.NewObjectID()

	result, err := collection.InsertOne(ctx, uvm)
	if err != nil {
		log.Printf("Could not create userdata: %v", err)
		return primitive.NilObjectID, err
	}

	oid := result.InsertedID.(primitive.ObjectID)

	return oid, err
}
