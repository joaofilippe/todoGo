package main

import (
	_ "github.com/lib/pq"

	"github.com/joaofilippe/todoGo/config"
	webserver "github.com/joaofilippe/todoGo/internal/adapters/web"
	"github.com/joaofilippe/todoGo/internal/application"
	di_connection "github.com/joaofilippe/todoGo/internal/di/connections"
	di_userusecases "github.com/joaofilippe/todoGo/internal/di/usecases"
	"github.com/joaofilippe/todoGo/internal/infra/user_service"
	userMigratons "github.com/joaofilippe/todoGo/internal/migrations/users_migrations"
	"github.com/joaofilippe/todoGo/pkg/logger"
)

func main() {
	logger := logger.NewLogger()
	appConfig := config.New(logger)

	connection := di_connection.GetConnection(logger, appConfig)

	if err := userMigratons.CreateUsersTable(connection); err != nil {
		logger.Logger.Error(err.Error())
	}

	//UseCases
	userUsecases := di_userusecases.NewUserUsecases(logger, appConfig)
	userService := user_service.NewUserService(
		userusecases.CreateUser,
		userUsecases.Login,
	)
	application := application.NewApplication(userService, logger)

	if err := webserver.NewServer(application).Run(); err != nil {
		logger.Logger.Error(err.Error())
	}
}
