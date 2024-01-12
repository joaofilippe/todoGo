package user

import userModels "github.com/joaofilippe/todoGo/application/models/user"

type IUserService interface {
	Create(user *userModels.NewUser) (int, error)
	Login(user *userModels.Login) (string, error)
}

type IUserUtils interface {
	validateNewUser(user *userModels.NewUser) error
	validateLogin(user *userModels.Login) error
	generateToken(user *userModels.User) (string, error)
}
