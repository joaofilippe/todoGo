package user

import (
	"github.com/google/uuid"
	usersModels "github.com/joaofilippe/todoGo/application/models/users"
)

// Repository represents the user repository
type IRepository interface {
	IReader
	IWriter
}

// IWriter represents the user repository writer
type IWriter interface {
	CreateUser(user usersModels.User) (uuid.UUID, error)
}

// IReader represents the user repository reader
type IReader interface {
	GetUserByUsername(username string) (usersModels.User, error)
	GetUserByEmail(email string) (usersModels.User, error)
}
