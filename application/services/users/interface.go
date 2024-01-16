package users

import (
	"github.com/google/uuid"
	usersModels "github.com/joaofilippe/todoGo/application/models/users"
)

type IUserService interface {
	Create(user *usersModels.NewUser) (uuid.UUID, error)
	Login(user *usersModels.Login) (string, error)
}

type IUserUtils interface {
	validateNewUser(user *usersModels.NewUser) error
	validateLogin(user *usersModels.Login) error
	generateToken(user *usersModels.User) (string, error)
}
