package user

import (
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

// CreateUserTable creates the user table
func (r *Repository) CreateUserTable() error {
	return r.Database.CreateUserTable()
}

// GetUserByUsername returns a user by username
func (r *Repository) GetUserByUsername(username string) (*userModels.User, error)

// GetUserByEmail returns a user by email
func (r *Repository) GetUserByEmail(email string) (*userModels.User, error)

// CreateUser creates a new user
func (r *Repository) CreateUser(user *userModels.User) (int, error)
