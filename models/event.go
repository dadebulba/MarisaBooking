package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        uuid.UUID               `json:"id"`
	Title     string                  `json:"title"`
	Metadata  *map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt time.Time               `json:"created_at"`
}
