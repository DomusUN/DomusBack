package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersMetadata struct {
	ID        primitive.ObjectID `bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	Name      string             `json:"name" bson:"name"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Direction string             `json:"direction" bson:"direction"`
	Phone     string             `json:"phone" bson:"phone"`
	Role      int                `json:"role" bson:"role"`
	//Client [] Array of ints
	//Worker [] Array of worker
}
