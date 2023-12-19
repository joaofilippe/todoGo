package user

import (
	"github.com/joaofilippe/todoGo/application/domain/models"
)
type Repository interface {
	CreateUser(user *models.User) (int, error)
}