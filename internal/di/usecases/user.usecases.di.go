package usecases

import user_usecases "github.com/joaofilippe/todoGo/internal/domain/usecases/user"

type UserUsecases struct {
	CreateUser *user_usecases.CreateUseCase
	Login      *user_usecases.LoginUsecase
}

func NewUserUsecases() *UserUsecases {
	return &UserUsecases{
		
	}

}