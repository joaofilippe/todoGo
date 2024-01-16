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

// FromDomain converts the User domain model to a UserDTO
func (u *UserDTO) FromDomain(user usersModels.User) {
	u.ID = user.ID
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Username = user.Username
	u.Email = user.Email
	u.Password = user.Password
	u.TodoIDs = user.TodoIDs
	u.BirthDate = user.BirthDate
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
	u.Active = user.Active
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

// FromDB converts the UserDB to a UserDTO
func (u *UserDTO) FromDB(user userDB.UserDB) {
	u.ID = user.ID
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Username = user.Username
	u.Email = user.Email
	u.Password = user.Password
	u.TodoIDs = user.TodoIDs
	u.BirthDate = user.BirthDate
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
	u.Active = user.Active
}