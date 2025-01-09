package routes

import (
	"github.com/dadebulba/marisabooking/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// User routes
	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)

	return router
}
