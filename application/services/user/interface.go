package user

import "github.com/joaofilippe/todoGo/application/domain/models"

type IUserService interface {
	Create(user *models.User) (int, error)
}

type IUserRepository interface {
	Create(todo *models.User) (int, error)
}