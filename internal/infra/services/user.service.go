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

func NewUserService(
	create userUsecases.CreateUseCase,
	login userUsecases.LoginUsecase,
) *UserService {
	return &UserService{}
}

func (s *UserService) Create(newUser user.NewUser) (uuid.UUID, error) {
	return s.createUsecase.Execute(newUser)
}

func (s *UserService) Login(login user.Login) (string, error) {
	return s.loginUsecase.Execute(login)
}

func (s *UserService) GetByID(id uuid.UUID) (user.User, error) {
	return user.User{}, nil
}

func (s *UserService) GetByEmail(email string) (user.User, error) {
	return user.User{}, nil
}

func (s *UserService) GetByUsername(username string) (user.User, error) {
	return user.User{}, nil
}
