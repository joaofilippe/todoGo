package user

import (
	"github.com/joaofilippe/todoGo/adapters/database/postgres"
	userModels "github.com/joaofilippe/todoGo/application/models/user"
)

// Repository represents the user repository
type Repository struct {
	Master *postgres.Connection
	Slave  *postgres.Connection
}

// NewUserRepository creates a new user repository
func NewUserRepository(writer *postgres.Connection, reader *postgres.Connection) *Repository {
	return &Repository{
		Master: writer,
		Slave:  reader,
	}
}

// GetUserByUsername returns a user by username
func (r *Repository) GetUserByUsername(username string) (*userModels.User, error)

// GetUserByEmail returns a user by email
func (r *Repository) GetUserByEmail(email string) (*userModels.User, error)

// CreateUser creates a new user
func (r *Repository) CreateUser(user *userModels.User) (int, error)
