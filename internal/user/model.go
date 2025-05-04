package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	RoleID       uuid.UUID `db:"role_id" json:"role_id"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
