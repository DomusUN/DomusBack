package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ServicesMetadata struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	ImageName string             `json:"image_name" bson:"image_name"`
	ServiceId int                `json:"service_id" bson:"service_id"`
}
