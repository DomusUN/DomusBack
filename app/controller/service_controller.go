package controller

import (
	"DomusBack/app/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServiceMetadata struct {
	Umr repository.ServideMetadata
}

func (uc ServiceMetadata) GetAllServices(c *gin.Context) {

	services, err := uc.Umr.GetAllServices()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, services)
	return
	
}
