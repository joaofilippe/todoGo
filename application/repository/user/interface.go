package user

import (
	"github.com/google/uuid"
	usersModels "github.com/joaofilippe/todoGo/application/models/users"
)

// Repository represents the user repository
type IRepository interface {
	IRepositoryReader
	IRepositoryWriter
}

// RepositoryWriter represents the user repository writer
type IRepositoryWriter interface {
	CreateUser(user *usersModels.User) (uuid.UUID, error)
}

// RepositoryReader represents the user repository reader
type IRepositoryReader interface {
	GetUserByUsername(username string) (*usersModels.User, error)
	GetUserByEmail(email string) (*usersModels.User, error)
}
