package services

import (
	"context"

	"github.com/dadebulba/marisabooking/models"
	"github.com/dadebulba/marisabooking/utils"
)

// GetAllGroups retrieves all groups from the database
func GetAllGroups() ([]models.Group, error) {
	query := "SELECT id, name, metadata, created_at FROM marisa.group"
	rows, err := utils.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var group models.Group
		if err := rows.Scan(&group.ID, &group.Name, &group.Metadata, &group.CreatedAt); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

// CreateGroup creates a new group in the database
func CreateGroup(group models.Group) (models.Group, error) {
	query := "INSERT INTO marisa.group (name, metadata) VALUES ($1, $2) RETURNING *"

	err := utils.Pool.QueryRow(
		context.Background(),
		query,
		group.Name,
		group.Metadata,
	).Scan(
		&group.ID,
		&group.Name,
		&group.CreatedAt,
		&group.Metadata,
	)
	if err != nil {
		return models.Group{}, err
	}
	return group, nil
}

// GetGroupByID retrieves a group by its ID from the database
func GetGroupByID(id string) (models.Group, error) {
	query := "SELECT id, name, metadata, created_at FROM marisa.group WHERE id = $1"
	var group models.Group
	err := utils.Pool.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(&group.ID, &group.Name, &group.Metadata, &group.CreatedAt)

	if err != nil {
		return models.Group{}, err
	}
	return group, nil
}

// UpdateGroup updates an existing group in the database
func UpdateGroup(id string, updatedGroup models.Group) (models.Group, error) {
	query := "UPDATE marisa.group SET name = COALESCE($1, name), metadata = COALESCE($2, metadata) WHERE id = $3 RETURNING *"
	err := utils.Pool.QueryRow(
		context.Background(),
		query,
		updatedGroup.Name,
		updatedGroup.Metadata,
		id,
	).Scan(&updatedGroup.ID, &updatedGroup.Name, &updatedGroup.CreatedAt, &updatedGroup.Metadata)
	if err != nil {
		return models.Group{}, err
	}
	return updatedGroup, nil
}

// DeleteGroup deletes a group by its ID from the database
func DeleteGroup(id string) error {
	query := "DELETE FROM marisa.group WHERE id = $1"
	_, err := utils.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
