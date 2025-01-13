package models

import "time"

// Group represents a group entity
type Group struct {
	ID        int                     `json:"id" gorm:"primary_key"`
	Name      string                  `json:"name" binding:"required"`
	Metadata  *map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt time.Time               `json:"created_at"`
}
