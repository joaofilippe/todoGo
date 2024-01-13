package user

import (
	"github.com/joaofilippe/todoGo/adapters/repository/postgres"
	userModels "github.com/joaofilippe/todoGo/application/models/user"
)

// UserRepository represents the user repository
type UserRepository struct {
	Master *postgres.Connection
	Slave  *postgres.Connection
}

// NewUserRepository creates a new user repository
func NewUserRepository(writer *postgres.Connection, reader *postgres.Connection) *UserRepository {
	return &UserRepository{
		Master: writer,
		Slave:  reader,
	}
}

// GetUserByUsername returns a user by username
func (r *UserRepository) GetUserByUsername(username string) (*userModels.User, error)

// GetUserByEmail returns a user by email
func (r *UserRepository) GetUserByEmail(email string) (*userModels.User, error)

// CreateUser creates a new user
func (r *UserRepository) CreateUser(user *userModels.User) (int, error)
