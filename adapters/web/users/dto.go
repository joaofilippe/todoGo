package users

import (
	"time"

	usersModel "github.com/joaofilippe/todoGo/application/models/users"
)

// NewUserDTO represents the new user DTO
type NewUserDTO struct {
	FirstName string
	LastName  string
	Username string
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
func (n *NewUserDTO) FromDTOToModel() (usersModel.NewUser, error) {
	
	newUser := usersModel.NewUser{
		FirstName: n.FirstName,
		LastName:  n.LastName,
		Username:  n.Username,
		Email:     n.Email,
		Password:  n.Password,
	}

	birthDate, err := time.Parse("2006-01-02", n.BirthDate)
	if err != nil {
		return usersModel.NewUser{}, err
	}

	newUser.BirthDate = birthDate

	return newUser, nil
}