package dto

import (
	"time"

	"github.com/google/uuid"

	userDB "github.com/joaofilippe/todoGo/internal/adapters/database/postgres/models"
	userEntity "github.com/joaofilippe/todoGo/internal/application/entities/user"
)

// UserDTO is the model for the user table
type NewUserDTO struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	BirthDate time.Time
}

// NewUserFromDomain is a function that creates a new UserDTO from Domain
func NewUserFromDomain(u userEntity.NewUser) *NewUserDTO {
	return &NewUserDTO{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		BirthDate: u.BirthDate,
	}
}

// ToDB converts the UserDTO to a NewUserDB
func (u *NewUserDTO) NewUserToDB() userDB.NewUserDB {
	return userDB.NewUserDB{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		BirthDate: u.BirthDate,
	}
}
