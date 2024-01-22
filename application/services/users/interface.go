package users

import (
	"github.com/google/uuid"
	usersModels "github.com/joaofilippe/todoGo/application/models/users"
)

type IService interface {
	Create(user usersModels.NewUser) (uuid.UUID, error)
	GetUserByID(id uuid.UUID) (usersModels.User, error)
	GetUserByUsername(username string) (usersModels.User, error)
	GetUserByEmail(email string) (usersModels.User, error)
	Login(user usersModels.Login) (string, error)
}

type IUtils interface {
	validateNewUser(user usersModels.NewUser) error
	validateLogin(user usersModels.Login) error
	generateToken(user usersModels.User) (string, error)
}
