package entities

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/pkg/email"
	"github.com/joaofilippe/todoGo/pkg/structs"
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

// Validate validates a new user
func (u *User) Validate() (bool, error) {
	fields, err := structs.CheckEmptyFields(&u)
	if err != nil {
		return false, err
	}

	if len(fields) > 0 {
		return false, errors.New("Empty fields: " + strings.Join(fields, ", "))
	}

	emailValid := email.Validate(u.Email)
	if !emailValid {
		return false, errors.New("invalid email")
	}

	return true, nil
}
