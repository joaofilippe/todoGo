package user

import (
	"github.com/google/uuid"
	usersModels "github.com/joaofilippe/todoGo/internal/application/models/users"
)

// IRepository represents the user repository
type IRepository interface {
	IReader
	IWriter
}

// IWriter represents the user repository writer
type IWriter interface {
	CreateNewUser(user usersModels.NewUser) error
}

// IReader represents the user repository reader
type IReader interface {
	GetUserByID(id uuid.UUID) (usersModels.User, error)
	GetUserByUsername(username string) (usersModels.User, error)
	GetUserByEmail(email string) (usersModels.User, error)
	CheckUserExists(username, email string) (bool, error) 
}
