package controller

import (
	"DomusBack/app/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ServiceMetadata struct {
	Mrs repository.ServiceMetadata
}

func (uc ServiceMetadata) GetAllServices(c *gin.Context) {

	services, err := uc.Mrs.GetAllServices()

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, services)
	return

}

func (uc ServiceMetadata) GetServiceById(c *gin.Context) {
	id := c.Param("id")
	idService, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	println("EXITO TOTAL")

	service, err := uc.Mrs.GetServiceById(idService)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	println("SUPER EXITO TOTAL")

	c.JSON(http.StatusAccepted, service)
	return

}
