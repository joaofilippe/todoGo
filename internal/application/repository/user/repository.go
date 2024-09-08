package user

import (
	"net/mail"

	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/adapters/database/postgres"
	userEntity "github.com/joaofilippe/todoGo/internal/application/entities/user"
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
func (r *Repository) GetUserByID(id uuid.UUID) (userEntity.User, error) {
	if id == uuid.Nil {
		return userEntity.User{}, ErrNoID
	}

	return r.Database.GetUserByID(id)
}

// GetUserByUsername returns a user by username
func (r *Repository) GetUserByUsername(username string) (userEntity.User, error) {
	if username == "" {
		return userEntity.User{}, ErrNoUsername
	}

	return r.Database.GetUserByUsername(username)
}

// GetUserByEmail returns a user by email
func (r *Repository) GetUserByEmail(email string) (userEntity.User, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return userEntity.User{}, ErrNoEmail
	}

	return r.Database.GetUserByEmail(email)
}

// CreateNewUser creates a new user
func (r *Repository) CreateNewUser(newUser userEntity.NewUser) error {
	return r.Database.CreateNewUser(newUser)
}

// CheckUserExists checks if a user exists
func (r *Repository) CheckUserExists(username, email string) (bool, error) {
	return r.Database.CheckUserExists(username, email)
}
