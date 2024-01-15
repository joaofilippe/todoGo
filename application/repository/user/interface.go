package user

import (
	"github.com/google/uuid"
	userModels "github.com/joaofilippe/todoGo/application/models/user"
)

// Repository represents the user repository
type IRepository interface {
	IRepositoryReader
	IRepositoryWriter
}

// RepositoryWriter represents the user repository writer
type IRepositoryWriter interface {
	CreateUser(user *userModels.User) (uuid.UUID, error)
}

// RepositoryReader represents the user repository reader
type IRepositoryReader interface {
	GetUserByUsername(username string) (*userModels.User, error)
	GetUserByEmail(email string) (*userModels.User, error)
}
