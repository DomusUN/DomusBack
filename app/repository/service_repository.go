package repository

import (
	"DomusBack/app/domain"
	"DomusBack/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type ServideMetadata interface {
	GetAllServices() ([]*domain.ServicesMetadata, error)
	GetServiceById(int) (*domain.ServicesMetadata, error)
}

type ServiceMetadataMongo struct{}

var collectionServices = database.OpenCollection(database.Client, "services")

func (u ServiceMetadataMongo) GetAllServices() ([]*domain.ServicesMetadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	var services []*domain.ServicesMetadata

	result, err := collectionServices.Find(ctx, bson.D{})

	if err != nil {
		log.Printf("Could not get the collections: %v", err)
		return []*domain.ServicesMetadata{}, err
	}

	for result.Next(ctx) {
		var service *domain.ServicesMetadata
		if err = result.Decode(&service); err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	return services, err
}

func (u ServiceMetadataMongo) GetServiceById(id int) (*domain.ServicesMetadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	var service *domain.ServicesMetadata

	err := collectionServices.FindOne(ctx, bson.M{"service_id": id}).Decode(&service)

	if err != nil {
		log.Printf("Could not get the collection: %v", err)
		return nil, err
	}

	return service, err
}
