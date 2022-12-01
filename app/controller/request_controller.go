package controller

import (
	"DomusBack/app/domain"
	"DomusBack/app/repository"
	"github.com/gin-gonic/gin"
	"net/http"
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
