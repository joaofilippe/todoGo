package userdatabase

import (
	"time"

	"github.com/google/uuid"

	userEntity "github.com/joaofilippe/todoGo/internal/domain/entities/user"
)

// NewUserDTO is the model for the user table
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
func NewUserFromDomain(u userEntity.User) *NewUserDTO {
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

// NewUserToDB converts the UserDTO to a NewUserDB
func (u *NewUserDTO) NewUserToDB() NewUserDB {
	return NewUserDB{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		BirthDate: u.BirthDate,
	}
}
