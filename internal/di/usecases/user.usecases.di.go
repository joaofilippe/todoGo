package di_userusecases

import (
	"github.com/joaofilippe/todoGo/config"
	"github.com/joaofilippe/todoGo/internal/di/repositories"
	"github.com/joaofilippe/todoGo/internal/domain/usecases/user"
	"github.com/joaofilippe/todoGo/pkg/logger"
)

type UserUsecases struct {
	CreateUser user_usecases.CreateUseCase
	Login      user_usecases.LoginUsecase
}

func NewUserUsecases(
	logger *logger.Logger,
	appConfig *config.App,
) *UserUsecases {
	userRepository := di_repositories.GetUserRepository(logger, appConfig)

	return &UserUsecases{
		CreateUser: user_usecases.NewCreateUserUseCase(userRepository),
		Login:      user_usecases.NewLoginUseCase(userRepository),
	}

}
