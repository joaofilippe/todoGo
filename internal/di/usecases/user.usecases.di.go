package userusecasesdi

import (
	"github.com/joaofilippe/todoGo/config"
	repositoriesdi "github.com/joaofilippe/todoGo/internal/di/repositories"
	"github.com/joaofilippe/todoGo/internal/domain/usecases/user"
	"github.com/joaofilippe/todoGo/pkg/logger"
)

type UserUsecases struct {
	CreateUser userusecases.CreateUseCase
	Login      userusecases.LoginUsecase
}

func NewUserUsecases(
	logger *logger.Logger,
	appConfig *config.App,
) *UserUsecases {
	userRepository := repositoriesdi.GetUserRepository(logger, appConfig)

	return &UserUsecases{
		CreateUser: userusecases.NewCreateUserUseCase(userRepository),
		Login:      userusecases.NewLoginUseCase(userRepository),
	}

}
