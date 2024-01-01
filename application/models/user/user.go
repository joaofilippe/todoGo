package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	TodoIDs   []uuid.UUID
	BirthDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Active    bool
}

type NewUser struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

type Login struct {
	Email    string
	Username string
	Password string
}
