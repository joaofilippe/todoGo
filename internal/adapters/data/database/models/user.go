package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// UserDB is the model for the user table
type UserDB struct {
	ID        uuid.UUID      `db:"id"`
	FirstName string         `db:"first_name"`
	LastName  string         `db:"last_name"`
	Username  string         `db:"username"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	TodoIDs   sql.NullString `db:"todo_ids"`
	BirthDate sql.NullString `db:"birthdate"`
	CreatedAt sql.NullString `db:"created_at"`
	UpdatedAt sql.NullString `db:"updated_at"`
	Active    bool           `db:"active"`
}

// NewUserDB is the model for the user table
type NewUserDB struct {
	ID        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	BirthDate time.Time `db:"birth_date"`
	CreatedAt time.Time `db:"created_at"`
	Active    bool      `db:"active"`
}
