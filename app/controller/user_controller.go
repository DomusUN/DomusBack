package controller

import (
	"DomusBack/app/domain"
	"DomusBack/app/repository"
	"github.com/gin-gonic/gin"
	"log"
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

	log.Printf("%v\n", user)
	oid, err := uc.Umr.Create(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"ID": oid})
}
