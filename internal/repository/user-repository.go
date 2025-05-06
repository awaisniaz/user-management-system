package repository

import (
	"context"

	"github.com/awaisniaz/user-management/internal/models"
	"github.com/awaisniaz/user-management/pkg/db"
)

func CreateUser(ctx context.Context, user models.User) error {
	query := `
		INSERT INTO users (id, email, password, role, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := db.DB.Exec(ctx, query,
		user.ID, user.Email, user.Password, user.Role, user.CreatedAt)

	return err
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, email, password, role, created_at FROM users WHERE email = $1`

	var user models.User
	err := db.DB.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
