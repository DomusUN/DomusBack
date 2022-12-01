package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestMetadata struct {
	ID           primitive.ObjectID `bson:"_id"`
	IdClient     primitive.ObjectID `json:"id_client" bson:"id_client"`
	IdWorker     primitive.ObjectID `json:"id_worker" bson:"id_worker"`
	RequestState int                `json:"state" bson:"state"`
	Date         string             `json:"date" bson:"date"`
}
