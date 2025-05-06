package models

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // do not expose in JSON
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
