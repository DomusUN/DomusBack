package app

import (
	"DomusBack/app/controller"
	"DomusBack/app/repository"
	"github.com/gin-gonic/gin"
)

var umc controller.UserMetadata

// TODO: Replace this with dependency injection.
func InitDepenedencies() {
	umr := repository.UserMetadataMongo{}
	umc = controller.UserMetadata{Umr: umr}
}

func InitRoutes(router *gin.Engine) error {
	router.POST("/users", umc.CreateUser)
	router.POST("/users/role/client/:id", umc.AddRoleClient)
	router.POST("/users/role/worker/:id", umc.AddRoleWorker)
	return nil
}
