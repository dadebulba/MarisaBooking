package services

import (
	"context"

	"github.com/dadebulba/marisabooking/models"
	"github.com/dadebulba/marisabooking/utils"
)

func GetAllUsers() ([]models.User, error) {
	rows, err := utils.Pool.Query(context.Background(), "SELECT id, email, phone, name, surname, birth, role, metadata, created_at, updated_at FROM marisa.user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Phone, &user.Name, &user.Surname, &user.Birth, &user.Role, &user.Metadata, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(user models.User) (models.User, error) {
	query := `
		INSERT INTO marisa.user (email, phone, name, surname, birth, role, metadata)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`
	err := utils.Pool.QueryRow(
		context.Background(),
		query,
		user.Email, user.Phone, user.Name, user.Surname, user.Birth, user.Role, user.Metadata,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return user, err
}
