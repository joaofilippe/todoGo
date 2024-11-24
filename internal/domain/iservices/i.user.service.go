package services

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/domain/entities"
)

type IUserService interface {
	Create(newUser entities.User) (uuid.UUID, error)
	Login(login entities.Login) (string, error)
	GetByID(id uuid.UUID) (entities.User, error)
	GetByEmail(email string) (entities.User, error)
	GetByUsername(username string) (entities.User, error)
}
