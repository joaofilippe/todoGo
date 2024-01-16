package users

import (
	usersModel "github.com/joaofilippe/todoGo/application/models/users"
)

// NewUserDTO represents the new user DTO
type NewUserDTO struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	BirthDate string
}

// FromRequestToDTO converts a request to a DTO
func (n *NewUserDTO) FromRequestToDTO(request *NewUserRequest) {
	n.FirstName = request.FirstName
	n.LastName = request.LastName
	n.Email = request.Email
	n.Password = request.Password
	n.BirthDate = request.BirthDate
}

// FromDTOToModel converts a DTO to a model
func (n *NewUserDTO) FromDTOToModel() *usersModel.NewUser {
	return &usersModel.NewUser{
		FirstName: n.FirstName,
		LastName:  n.LastName,
		Email:     n.Email,
		Password:  n.Password,
	}
}