package dto

import (
	"time"

	"github.com/google/uuid"

	usersModels "github.com/joaofilippe/todoGo/application/models/users"
	userDB "github.com/joaofilippe/todoGo/adapters/database/postgres/models"
)

// UserDTO is the model for the user table
type UserDTO struct {
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

// UserFromDomain is a function that creates a new UserDTO
func UserFromDomain(u usersModels.User) *UserDTO {
	return &UserDTO{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		TodoIDs:   u.TodoIDs,
		BirthDate: u.BirthDate,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Active:    u.Active,
	}
}

// UserFromDB converts the UserDB to a UserDTO
func UserFromDB(user userDB.UserDB) *UserDTO{
	return &UserDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		TodoIDs:   user.TodoIDs,
		BirthDate: user.BirthDate,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Active:    user.Active,
	}
}

// ToDomain converts the UserDTO to a User domain model
func (u *UserDTO) ToDomain() usersModels.User {
	return usersModels.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		TodoIDs:   u.TodoIDs,
		BirthDate: u.BirthDate,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Active:    u.Active,
	}
}

// ToDB converts the UserDTO to a UserDB
func (u *UserDTO) ToDB() userDB.UserDB {
	return userDB.UserDB{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		TodoIDs:   u.TodoIDs,
		BirthDate: u.BirthDate,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Active:    u.Active,
	}
}