package app

import (
	"DomusBack/app/controller"
	"DomusBack/app/repository"
	"github.com/gin-gonic/gin"
)

var umc controller.UserMetadata

var umcs controller.ServiceMetadata

// TODO: Replace this with dependency injection.
func InitDepenedencies() {
	umr := repository.UserMetadataMongo{}
	umc = controller.UserMetadata{Umr: umr}

	umrs := repository.ServiceMetadataMongo{}
	umcs = controller.ServiceMetadata{Umr: umrs}
}

func InitRoutes(router *gin.Engine) error {
	// Create User
	router.POST("/users", umc.CreateUser)
	// Add roles
	router.POST("/users/role/client/:id", umc.AddRoleClient)
	router.POST("/users/role/worker/:id", umc.AddRoleWorker)

	// Get services
	router.GET("/services", umcs.GetAllServices)
	router.GET("/services/:id", umcs.GetServiceById)

	return nil
}
