package user

import (
	"github.com/google/uuid"
	userEntity "github.com/joaofilippe/todoGo/internal/domain/entities/user"
)

// IRepository represents the user repository
type IRepository interface {
	IReader
	IWriter
}

// IWriter represents the user repository writer
type IWriter interface {
	CreateNewUser(user userEntity.NewUser) error
}

// IReader represents the user repository reader
type IReader interface {
	GetUserByID(id uuid.UUID) (userEntity.User, error)
	GetUserByUsername(username string) (userEntity.User, error)
	GetUserByEmail(email string) (userEntity.User, error)
	CheckUserExists(username, email string) (bool, error)
}
