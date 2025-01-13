package models

import "time"

type User struct {
	ID        int                     `json:"id"`
	Email     string                  `json:"email"`
	Phone     string                  `json:"phone,omitempty"`
	Name      string                  `json:"name,omitempty"`
	Surname   string                  `json:"surname,omitempty"`
	Birth     time.Time               `json:"birth,omitempty"`
	Role      string                  `json:"role"`
	Metadata  *map[string]interface{} `json:"metadata,omitempty"` // Store as JSON string
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
}
