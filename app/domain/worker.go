package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ServicesWorker struct {
	ServiceID primitive.ObjectID `json:"service_id" bson:"service_id"`
	Score     float32            `json:"score" bson:"score"`
	Price     int32              `json:"price"  bson:"price"`
}

type WorkerMetadata struct {
	NumberId    string           `json:"number_id" bson:"number_id"`
	Description string           `json:"description" bson:"description"`
	Score       float32          `json:"score" bson:"score"`
	Places      []int            `json:"places" bson:"places"`
	Reviews     []string         `json:"reviews" bson:"reviews"`
	ArrServices []ServicesWorker `json:"arr_services" bson:"arr_services"`
}
