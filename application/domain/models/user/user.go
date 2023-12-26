package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	BirthDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Active   bool
}

type NewUser struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}