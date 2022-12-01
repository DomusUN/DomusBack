package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ClientMetadata struct {
	Historyservices []primitive.ObjectID `json:"h_services" bson:"h_services"`
}
