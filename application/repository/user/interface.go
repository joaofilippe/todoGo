package user

import (
	userModels "github.com/joaofilippe/todoGo/application/domain/models/user"
)
type Repository interface {
	GetUserByEmail(email string) (*userModels.User, error)
	GetUserByUsername(username string) (*userModels.User, error)
	CreateUser(user *userModels.User) (int, error)
}