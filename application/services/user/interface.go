package user

import (
	"github.com/google/uuid"
	userModels "github.com/joaofilippe/todoGo/application/models/user"
)

type IUserService interface {
	Create(user *userModels.NewUser) (uuid.UUID, error)
	Login(user *userModels.Login) (string, error)
}

type IUserUtils interface {
	validateNewUser(user *userModels.NewUser) error
	validateLogin(user *userModels.Login) error
	generateToken(user *userModels.User) (string, error)
}
