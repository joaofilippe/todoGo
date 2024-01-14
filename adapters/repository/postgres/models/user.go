package models

import (
	"time"

	"github.com/google/uuid"
)

// UserDB is the model for the user table
type UserDB struct {
	ID        uuid.UUID   `db:"id"`
	FirstName string      `db:"first_name"`
	LastName  string      `db:"last_name"`
	Username  string      `db:"username"`
	Email     string      `db:"email"`
	Password  string      `db:"password"`
	TodoIDs   []uuid.UUID `db:"todo_ids"`
	BirthDate time.Time   `db:"birth_date"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt time.Time   `db:"updated_at"`
	Active    bool        `db:"active"`
}
