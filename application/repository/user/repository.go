package user

import (
	"net/mail"

	"github.com/joaofilippe/todoGo/adapters/database/postgres"
	usersModels "github.com/joaofilippe/todoGo/application/models/users"
)

// Repository represents the user repository
type Repository struct {
	Database *postgres.Database
}

// NewUserRepository creates a new user repository
func NewUserRepository(writer *postgres.Connection, reader *postgres.Connection) *Repository {
	database := postgres.NewDatabase(writer, reader)

	return &Repository{
		Database: database,
	}
}

// GetUserByUsername returns a user by username
func (r *Repository) GetUserByUsername(username string) (usersModels.User, error) {
	return usersModels.User{}, nil
}

// GetUserByEmail returns a user by email
func (r *Repository) GetUserByEmail(email string) (usersModels.User, error){
	_, err := mail.ParseAddress(email)
	if err != nil {
		return usersModels.User{}, err
	}


	return usersModels.User{}, nil
}

// CreateUser creates a new user
func (r *Repository) CreateUser(newUser usersModels.NewUser) error{
	return r.Database.CreateUser(newUser)
}
