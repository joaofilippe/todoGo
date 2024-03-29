package user

import (
	"net/mail"

	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/adapters/database/postgres"
	usersModels "github.com/joaofilippe/todoGo/internal/application/models/users"
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

// GetUserByID returns a user by id
func (r *Repository) GetUserByID(id uuid.UUID) (usersModels.User, error) {
	if id == uuid.Nil {
		return usersModels.User{}, ErrNoID
	}

	return r.Database.GetUserByID(id)
}

// GetUserByUsername returns a user by username
func (r *Repository) GetUserByUsername(username string) (usersModels.User, error) {
	if username == "" {
		return usersModels.User{}, ErrNoUsername
	}

	return r.Database.GetUserByUsername(username)
}

// GetUserByEmail returns a user by email
func (r *Repository) GetUserByEmail(email string) (usersModels.User, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return usersModels.User{}, ErrNoEmail
	}

	return r.Database.GetUserByEmail(email)
}

// CreateNewUser creates a new user
func (r *Repository) CreateNewUser(newUser usersModels.NewUser) error {
	return r.Database.CreateNewUser(newUser)
}

// CheckUserExists checks if a user exists
func (r *Repository) CheckUserExists(username, email string) (bool, error) {
	return r.Database.CheckUserExists(username, email)
}
