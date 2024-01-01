package user

import userModels "github.com/joaofilippe/todoGo/application/domain/models/user"

type IUserService interface {
	Create(user *userModels.NewUser) (int, error)
	Login(user *userModels.Login) (string, error)
}

type IUserUtils interface {
	ValidateLogin(user *userModels.Login) error
	ValidateUser(user *userModels.NewUser) error
	GenerateToken(user *userModels.User) (string, error)
}
