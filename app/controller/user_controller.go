package controller

import (
	"DomusBack/app/domain"
	"DomusBack/app/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// TODO: Replace with proper dependency injection.
type UserMetadata struct {
	Umr repository.UserMetadata
}

func (uc UserMetadata) CreateUser(c *gin.Context) {
	var user domain.UsersMetadata

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	oid, err := uc.Umr.Create(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"ID": oid})
}

// Add role client
func (uc UserMetadata) AddRoleClient(c *gin.Context) {
	var client *domain.ClientMetadata
	id := c.Param("id")
	// Check ID
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check the BODY of the request
	if err := c.ShouldBindJSON(&client); err != nil {
		// If it doesnt bind
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Add client role
	// Check histories  list
	if len(client.Historyservices) == 0 {
		client.Historyservices = []primitive.ObjectID{}
	}
	_, err = uc.Umr.AddRoleClient(userID, client)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"ID": id})
	return
}

// Add role worker
func (uc UserMetadata) AddRoleWorker(c *gin.Context) {
	var worker *domain.WorkerMetadata
	id := c.Param("id")
	// Check ID
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check the BODY of the request
	if err := c.ShouldBindJSON(&worker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Add worker role
	// Check places list
	if len(worker.Places) == 0 {
		worker.Places = []int{}
	}
	if len(worker.Reviews) == 0 {
		worker.Reviews = []string{}
	}
	// Check the services list
	if len(worker.ArrServices) == 0 {
		worker.ArrServices = []domain.ServicesWorker{}
	}
	_, err = uc.Umr.AddRoleWorker(userID, worker)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"ID": id})
	return
}

func (uc UserMetadata) GetAllWorkers(c *gin.Context) {
	
	return
}
