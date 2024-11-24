package userusecasesdi

import (
	"github.com/joaofilippe/todoGo/config"
	repositoriesdi "github.com/joaofilippe/todoGo/internal/di/repositories"
	"github.com/joaofilippe/todoGo/internal/domain/usecases/user"
	"github.com/joaofilippe/todoGo/pkg/logger"
)

// UserUsecases contains use cases related to user operations.
type UserUsecases struct {
	Create userusecases.CreateUseCase
	Login      userusecases.LoginUsecase
}

// NewUserUsecases creates a new instance of UserUsecases.
func NewUserUsecases(
	logger *logger.Logger,
	appConfig *config.App,
) *UserUsecases {
	userRepository := repositoriesdi.GetUserRepository(logger, appConfig)

	return &UserUsecases{
		Create: userusecases.NewCreateUseCase(userRepository),
		Login:      userusecases.NewLoginUseCase(userRepository),
	}

}
