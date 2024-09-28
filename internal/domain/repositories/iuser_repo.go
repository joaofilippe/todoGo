package userRepo

import (
	"github.com/google/uuid"
	userEntity "github.com/joaofilippe/todoGo/internal/domain/entities/user"
)

type IUserRepo struct {
	Writer IUserWriter
	Reader IUserReader
}

type IUserWriter interface {
	Create(user userEntity.NewUser) (uuid.UUID, error)
}

type IUserReader interface {
	GetUserByEmail(email string)(*userEntity.User, error)
}
