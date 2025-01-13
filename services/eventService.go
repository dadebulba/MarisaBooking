package services

import (
	"context"

	"github.com/dadebulba/marisabooking/models"
	"github.com/dadebulba/marisabooking/utils"
)

func GetAllEvent() ([]models.Event, error) {
	query := `SELECT id, title, metadata, created_at FROM marisa.event`
	rows, err := utils.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.ID, &event.Title, &event.Metadata, &event.CreatedAt); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func CreateEvent(event models.Event) (models.Event, error) {
	query := `
		INSERT INTO marisa.event (title, metadata)
		VALUES ($1, $2)
		RETURNING *;
	`
	err := utils.Pool.QueryRow(
		context.Background(),
		query,
		event.Title, event.Metadata,
	).Scan(&event.ID, &event.Title, &event.Metadata, &event.CreatedAt)

	return event, err
}
