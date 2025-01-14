package controllers

import (
	"net/http"

	"github.com/dadebulba/marisabooking/models"
	"github.com/dadebulba/marisabooking/services"

	"github.com/gin-gonic/gin"
)

// GetEvents retrieves all events
func GetEvents(c *gin.Context) {
	events, err := services.GetAllEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

// CreateEvent creates a new event
func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	createdEvent, err := services.CreateEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdEvent)
}
