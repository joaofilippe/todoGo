package user

import (
	userRepo "github.com/joaofilippe/todoGo/internal/application/repository/user"
	userEntity "github.com/joaofilippe/todoGo/internal/application/entities/user"
)

type CreateUserUsecase struct {
	repository userRepo.IRepository
}

func (u *CreateUserUsecase) Execute (newUser userEntity.NewUser) error {
	return u.repository.CreateNewUser(newUser)
}