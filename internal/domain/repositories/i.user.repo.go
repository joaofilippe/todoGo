package irepositories

import (
	"github.com/google/uuid"
	userEntity "github.com/joaofilippe/todoGo/internal/domain/entities/user"
)

type IUserRepo interface {
	CreateNewUser(newUser userEntity.User) (uuid.UUID, error)
	GetUserByID(id uuid.UUID) (userEntity.User, error)
	GetUserByEmail(email string) (userEntity.User, error)
	GetUserByUsername(username string) (userEntity.User, error)
}

