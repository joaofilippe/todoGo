package users

import (
	"time"

	"github.com/joaofilippe/todoGo/internal/domain/entities"
)

// NewUserDTO represents the new user DTO
type NewUserDTO struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	BirthDate string
}

// FromRequestToDTO converts a request to a DTO
func (n *NewUserDTO) FromRequestToDTO(request NewUserRequest) {
	n.FirstName = request.FirstName
	n.LastName = request.LastName
	n.Username = request.Username
	n.Email = request.Email
	n.Password = request.Password
	n.BirthDate = request.BirthDate
}

// FromDTOToModel converts a DTO to a model
func (n *NewUserDTO) FromDTOToModel() (entities.User, error) {

	newUser := entities.User{
		FirstName: n.FirstName,
		LastName:  n.LastName,
		Username:  n.Username,
		Email:     n.Email,
		Password:  n.Password,
	}

	birthDate, err := time.Parse("2006-01-02", n.BirthDate)
	if err != nil {
		return entities.User{}, err
	}

	newUser.BirthDate = birthDate

	return newUser, nil
}
