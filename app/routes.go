package app

import (
	"DomusBack/app/controller"
	"DomusBack/app/repository"
	"github.com/gin-gonic/gin"
)

var umc controller.UserMetadata

var ucs controller.ServiceMetadata

var ucr controller.RequestMetadata

// TODO: Replace this with dependency injection.
func InitDepenedencies() {
	// User
	umr := repository.UserMetadataMongo{}
	umc = controller.UserMetadata{Umr: umr}

	// Service
	mrs := repository.ServiceMetadataMongo{}
	ucs = controller.ServiceMetadata{Mrs: mrs}

	// Request
	rmr := repository.RequestMetadataMongo{}
	ucr = controller.RequestMetadata{Rmr: rmr}

	// Workers
}

func InitRoutes(router *gin.Engine) error {
	// Create User
	router.POST("/users", umc.CreateUser)
	// Add roles
	router.POST("/users/role/client/:id", umc.AddRoleClient)
	router.POST("/users/role/worker/:id", umc.AddRoleWorker)

	// Get services
	router.GET("/services", ucs.GetAllServices)
	router.GET("/services/:id", ucs.GetServiceById)

	// Create Request
	router.POST("/requests", ucr.Create)
	router.PUT("/requests/:id/state/:state", ucr.ChangeState)

	//Get requests
	router.GET("/requests", ucr.GetAllRequests)
	router.GET("/requests/worker/:id", ucr.GetRequestByWorker)

	// Get workers
	router.GET("/users/workers", umc.GetAllWorkers)
	router.GET("/users/workers/query", umc.GetWorkersByService)

	return nil
}
