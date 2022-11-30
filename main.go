package main

import (
	"DomusBack/app"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	app.InitDepenedencies() // MUST initializate dependencies, otherwise you'll get like 1M errors.

	err := app.InitRoutes(router)
	if err != nil {
		return
	}

	errRouter := router.Run()
	if errRouter != nil {
		return
	}
}
