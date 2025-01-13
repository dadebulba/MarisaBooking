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

	//Event routes
	router.GET("/events", controllers.GetEvents)
	router.POST("/events", controllers.CreateEvent)

	// Event-item routes
	router.GET("/event-items", controllers.GetAllEventItems)
	router.POST("/event-items", controllers.CreateEventItem)
	router.GET("/event-items/:id", controllers.GetEventItem)
	router.PUT("/event-items/:id", controllers.UpdateEventItem)
	router.DELETE("/event-items/:id", controllers.DeleteEventItem)

	// Group routes
	router.GET("/groups", controllers.GetGroups)
	router.POST("/groups", controllers.CreateGroup)
	router.GET("/groups/:id", controllers.GetGroup)
	router.PUT("/groups/:id", controllers.UpdateGroup)
	router.DELETE("/groups/:id", controllers.DeleteGroup)

	return router
}
