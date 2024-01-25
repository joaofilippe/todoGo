package users

import (
	"errors"
	"github.com/joaofilippe/todoGo/pkg/structs"
	"strings"
)

// Login represents a login request
type Login struct {
	Email    string
	Username string
	Password string
}

// Validate validates a new user
func (l Login) Validate() error {
	fields, err := structs.CheckEmptyFields(l)
	if err != nil {
		return err
	}

	if len(fields) > 2 {
		return errors.New("Empty fields: " + strings.Join(fields, ", "))
	}

	return nil
}
