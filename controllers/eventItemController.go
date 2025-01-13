package controllers

import (
	"net/http"
	"strconv"

	"github.com/dadebulba/marisabooking/models"
	"github.com/dadebulba/marisabooking/services"

	"github.com/gin-gonic/gin"
)

// GetAllEventItems retrieves all event items
func GetAllEventItems(c *gin.Context) {
	eventItems, err := services.GetAllEventItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eventItems)
}

// CreateEventItem creates a new event item
func CreateEventItem(c *gin.Context) {
	var eventItem models.EventItem
	if err := c.ShouldBindJSON(&eventItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	eventItem, err := services.CreateEventItem(eventItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, eventItem)
}

// GetEventItem retrieves a event item by ID
func GetEventItem(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	eventItem, err := services.GetEventItemByID(idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event item not found"})
		return
	}
	c.JSON(http.StatusOK, eventItem)
}

// UpdateEventItem updates an existing event item
func UpdateEventItem(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var eventItem models.EventItem
	if err := c.ShouldBindJSON(&eventItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	eventItem, err = services.UpdateEventItem(idInt, eventItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eventItem)
}

// DeleteEventItem deletes a event item by ID
func DeleteEventItem(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := services.DeleteEventItem(idInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
