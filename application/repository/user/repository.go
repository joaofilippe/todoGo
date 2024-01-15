package user

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/adapters/database/postgres"
	userModels "github.com/joaofilippe/todoGo/application/models/user"
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
func (r *Repository) GetUserByUsername(username string) (*userModels.User, error) {
	return nil, nil
}

// GetUserByEmail returns a user by email
func (r *Repository) GetUserByEmail(email string) (*userModels.User, error){
	return nil, nil
}

// CreateUser creates a new user
func (r *Repository) CreateUser(user *userModels.User) (uuid.UUID, error){
	return uuid.UUID{}, nil
}
