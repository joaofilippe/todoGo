package data

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/adapters/data/database"
 	userEntity "github.com/joaofilippe/todoGo/internal/domain/entities/user"
)

// UserRepository represents the user repository
type UserRepository struct {
	Writer database.IUserWriter
	Reader database.IUserReader
}

func (u *UserRepository) CreateNewUser(newUser userEntity.User) (uuid.UUID, error) {
	panic("TODO: Implement")
}

func (u *UserRepository) GetUserByID(id uuid.UUID) (userEntity.User, error) {
	panic("TODO: Implement")
}

func (u *UserRepository) GetUserByEmail(email string) (userEntity.User, error) {
	panic("TODO: Implement")
}

func (u *UserRepository) GetUserByUsername(username string) (userEntity.User, error) {
	panic("TODO: Implement")
}
