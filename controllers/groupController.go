package controllers

import (
	"net/http"

	"github.com/dadebulba/marisabooking/models"
	"github.com/dadebulba/marisabooking/services"

	"github.com/gin-gonic/gin"
)

// GetGroups retrieves all groups
func GetGroups(c *gin.Context) {
	groups, err := services.GetAllGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groups)
}

// CreateGroup creates a new group
func CreateGroup(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	group, err := services.CreateGroup(group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, group)
}

// GetGroup retrieves a group by ID
func GetGroup(c *gin.Context) {
	id := c.Param("id")
	group, err := services.GetGroupByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}
	c.JSON(http.StatusOK, group)
}

// UpdateGroup updates an existing group
func UpdateGroup(c *gin.Context) {
	id := c.Param("id")
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	group, err := services.UpdateGroup(id, group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, group)
}

// DeleteGroup deletes a group by ID
func DeleteGroup(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteGroup(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
