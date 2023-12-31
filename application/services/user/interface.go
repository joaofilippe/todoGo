package user

import userModels "github.com/joaofilippe/todoGo/application/domain/models/user"

type IUserService interface {
	Create(user *userModels.NewUser) (int, error)
}