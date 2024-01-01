package user

import userModels "github.com/joaofilippe/todoGo/application/models/user"

type IUserService interface {
	Create(user *userModels.NewUser) (int, error)
	Login(user *userModels.Login) (string, error)
}

type IUserUtils interface {
	validateLogin(user *userModels.Login) error
	validateUser(user *userModels.NewUser) error
	generateToken(user *userModels.User) (string, error)
}
