package services

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/domain/entities/user"
	userUsecases "github.com/joaofilippe/todoGo/internal/domain/usecases/user"
)

type UserService struct {
	createUsecase userUsecases.CreateUseCase
	loginUsecase  userUsecases.LoginUsecase
}

func NewUserService()

func (s *UserService) Create(newUser user.NewUser) (uuid.UUID, error) {
	return s.createUsecase.Execute(newUser)
}

func (s *UserService) Login(login user.Login) (string, error) {
	return s.loginUsecase.Execute(login)
}
