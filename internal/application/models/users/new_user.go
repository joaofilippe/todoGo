package users

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/pkg/email"
	"github.com/joaofilippe/todoGo/pkg/structs"
)

// NewUser represents a new user
type NewUser struct {
	ID 	  uuid.UUID
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	BirthDate time.Time
}

// Validate validates a new user
func (n NewUser) Validate() (bool, error) {
	fields, err := structs.CheckEmptyFields(n)
	if err != nil {
		return false, err
	}

	if len(fields) > 0 {
		return false, errors.New("Empty fields: " + strings.Join(fields, ", "))
	}

	emailValid:= email.Validate(n.Email)
	if !emailValid {
		return false, errors.New("Invalid email")
	}

	return true, nil
}