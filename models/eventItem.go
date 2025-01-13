package models

import (
	"time"

	"github.com/google/uuid"
)

type ItemKind string

const (
	KindConference ItemKind = "document"
	KindWorkshop   ItemKind = "payment"
)

type EventItem struct {
	ID               int       `json:"id" gorm:"primary_key"`
	EventID          uuid.UUID `json:"event_id"`
	Name             string    `json:"title"`
	PaymentLink      string    `json:"payment_link"`
	DocumentTamplate uuid.UUID `json:"document_template"`
	Kind             ItemKind  `json:"kind"`
	ExpiresAt        time.Time `json:"expires_at"`
	CreatedAt        time.Time `json:"created_at"`
}
