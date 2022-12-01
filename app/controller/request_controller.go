package controller

import (
	"DomusBack/app/domain"
	"DomusBack/app/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
)

type RequestMetadata struct {
	Rmr repository.RequestMetadata
}

func (uc RequestMetadata) Create(c *gin.Context) {
	var request domain.RequestMetadata

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	oid, err := uc.Rmr.Create(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"request_id": oid})

	return
}

func (uc RequestMetadata) ChangeState(c *gin.Context) {
	id := c.Param("id")
	state := c.Param("state")

	// Check ID
	idRequest, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stateRequest, err := strconv.Atoi(state)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	oid, err := uc.Rmr.ChangeState(idRequest, stateRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"request_id": oid})

	return
}
