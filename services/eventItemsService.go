package services

import (
	"context"

	"github.com/dadebulba/marisabooking/models"
	"github.com/dadebulba/marisabooking/utils"
)

// GetAllEventItems retrieves all event items from the database
func GetAllEventItems() ([]models.EventItem, error) {
	query := "SELECT id, event_id, name, payment_link, document_template, kind, expires_at, created_at FROM marisa.event_item"
	rows, err := utils.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventItems []models.EventItem
	for rows.Next() {
		var eventItem models.EventItem
		if err := rows.Scan(&eventItem.ID, &eventItem.EventID, &eventItem.Name, &eventItem.PaymentLink, &eventItem.DocumentTamplate, &eventItem.Kind, &eventItem.ExpiresAt, &eventItem.CreatedAt); err != nil {
			return nil, err
		}
		eventItems = append(eventItems, eventItem)
	}
	return eventItems, nil
}

// CreateEventItem creates a new event item in the database
func CreateEventItem(eventItem models.EventItem) (models.EventItem, error) {
	query := "INSERT INTO marisa.event_item (event_id, name, payment_link, document_template, kind, expires_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at"

	err := utils.Pool.QueryRow(
		context.Background(),
		query,
		eventItem.EventID,
		eventItem.Name,
		eventItem.PaymentLink,
		eventItem.DocumentTamplate,
		eventItem.Kind,
		eventItem.ExpiresAt,
	).Scan(
		&eventItem.ID,
		&eventItem.CreatedAt,
	)
	if err != nil {
		return models.EventItem{}, err
	}
	return eventItem, nil
}

// GetEventItemByID retrieves an event item by its ID from the database
func GetEventItemByID(id int) (models.EventItem, error) {
	query := "SELECT id, event_id, name, payment_link, document_template, kind, expires_at, created_at FROM marisa.event_item WHERE id = $1"
	var eventItem models.EventItem
	err := utils.Pool.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(&eventItem.ID, &eventItem.EventID, &eventItem.Name, &eventItem.PaymentLink, &eventItem.DocumentTamplate, &eventItem.Kind, &eventItem.ExpiresAt, &eventItem.CreatedAt)

	if err != nil {
		return models.EventItem{}, err
	}
	return eventItem, nil
}

// UpdateEventItem updates an existing event item in the database
func UpdateEventItem(id int, updatedEventItem models.EventItem) (models.EventItem, error) {
	query := "UPDATE marisa.event_item SET event_id = COALESCE($1, event_id), name = COALESCE($2, name), payment_link = COALESCE($3, payment_link), document_template = COALESCE($4, document_template), kind = COALESCE($5, kind), expires_at = COALESCE($6, expires_at) WHERE id = $7 RETURNING id, event_id, name, payment_link, document_template, kind, expires_at, created_at"
	err := utils.Pool.QueryRow(
		context.Background(),
		query,
		updatedEventItem.EventID,
		updatedEventItem.Name,
		updatedEventItem.PaymentLink,
		updatedEventItem.DocumentTamplate,
		updatedEventItem.Kind,
		updatedEventItem.ExpiresAt,
		id,
	).Scan(&updatedEventItem.ID, &updatedEventItem.EventID, &updatedEventItem.Name, &updatedEventItem.PaymentLink, &updatedEventItem.DocumentTamplate, &updatedEventItem.Kind, &updatedEventItem.ExpiresAt, &updatedEventItem.CreatedAt)
	if err != nil {
		return models.EventItem{}, err
	}
	return updatedEventItem, nil
}

// DeleteEventItem deletes an event item by its ID from the database
func DeleteEventItem(id int) error {
	query := "DELETE FROM marisa.event_item WHERE id = $1"
	_, err := utils.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
