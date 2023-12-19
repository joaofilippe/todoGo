package user

import (
	"github.com/joaofilippe/todoGo/application/models"
)
type Repository interface {
	CreateUser(user *models.User) (int, err)
}