package data

import (
	"github.com/google/uuid"
	userEntity "github.com/joaofilippe/todoGo/internal/domain/entities/user"
	irepositories "github.com/joaofilippe/todoGo/internal/domain/repositories"
)

// UserRepository represents the user repository
type UserRepository struct {
	Writer irepositories.IUserWriter
	Reader irepositories.IUserReader
}

// CreateNewUser creates a new user in the repository
func (u *UserRepository) CreateNewUser(newUser userEntity.User) (uuid.UUID, error) {
	return u.Writer.CreateNewUser(newUser)
}

// GetUserByID retrieves a user by their ID from the repository
func (u *UserRepository) GetUserByID(id uuid.UUID) (userEntity.User, error) {
	return u.Reader.GetUserByID(id)
}

// GetUserByEmail retrieves a user by their email from the repository
func (u *UserRepository) GetUserByEmail(email string) (userEntity.User, error) {
	return u.Reader.GetUserByEmail(email)
}

// GetUserByUsername retrieves a user by their username from the repository
func (u *UserRepository) GetUserByUsername(username string) (userEntity.User, error) {
	return u.Reader.GetUserByUsername(username)
}
