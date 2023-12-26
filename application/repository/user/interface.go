package user

import (
	userModels "github.com/joaofilippe/todoGo/application/domain/models/user"
)
type Repository interface {
	CreateUser(user *userModels.User) (int, error)
}