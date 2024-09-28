package services

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/domain/entities/user"
 )

type IUserService interface {
	Create(newUser user.NewUser) (uuid.UUID, error)
	Login(login user.Login)(string, error)
}
