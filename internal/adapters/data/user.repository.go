package data

import (
	"github.com/google/uuid"
	userdatabase "github.com/joaofilippe/todoGo/internal/adapters/data/database/user"
	"github.com/joaofilippe/todoGo/internal/domain/entities"
)

// UserRepository represents the user repository
type UserRepository struct {
	Writer *userdatabase.Writer
	Reader *userdatabase.Reader
}

// CreateNewUser creates a new user in the repository
func (u *UserRepository) CreateNewUser(newUser entities.User) (uuid.UUID, error) {
	return u.Writer.CreateNewUser(newUser)
}

// GetUserByID retrieves a user by their ID from the repository
func (u *UserRepository) GetUserByID(id uuid.UUID) (entities.User, error) {
	return u.Reader.GetUserByID(id)
}

// GetUserByEmail retrieves a user by their email from the repository
func (u *UserRepository) GetUserByEmail(email string) (entities.User, error) {
	return u.Reader.GetUserByEmail(email)
}

// GetUserByUsername retrieves a user by their username from the repository
func (u *UserRepository) GetUserByUsername(username string) (entities.User, error) {
	return u.Reader.GetUserByUsername(username)
}
