package users

import (
	"github.com/google/uuid"
	userEntity "github.com/joaofilippe/todoGo/internal/application/entities/user"
)

type IService interface {
	Create(user userEntity.NewUser) (uuid.UUID, error)
	GetUserByID(id uuid.UUID) (userEntity.User, error)
	GetUserByUsername(username string) (userEntity.User, error)
	GetUserByEmail(email string) (userEntity.User, error)
	Login(user userEntity.Login) (string, error)
}

type IUtils interface {
	validateNewUser(user userEntity.NewUser) error
	validateLogin(user userEntity.Login) error
	generateToken(user userEntity.User) (string, error)
}
