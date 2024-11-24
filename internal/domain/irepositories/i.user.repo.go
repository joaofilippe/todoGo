package irepositories

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/domain/entities"
)

// IUserRepo defines the repository interface for user operations.
type IUserRepo interface {
	CreateNewUser(newUser entities.User) (uuid.UUID, error)
	GetUserByID(id uuid.UUID) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	GetUserByUsername(username string) (entities.User, error)
}
