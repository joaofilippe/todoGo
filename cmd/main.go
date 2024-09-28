package main

import (
	_ "github.com/lib/pq"

	"github.com/joaofilippe/todoGo/config"
 	"github.com/joaofilippe/todoGo/internal/adapters/web"
	"github.com/joaofilippe/todoGo/internal/application"
 	userMigratons "github.com/joaofilippe/todoGo/internal/migrations/users"
 	"github.com/joaofilippe/todoGo/pkg/logger"
)

func main() {
	logger := logger.NewLogger()
	appConfig := config.NewApp(*logger)


	if err := userMigratons.CreateUsersTable(masterConnectionDB); err != nil {
		logger.Logger.Error(err.Error())
	}


	//UseCases


	userService := users.NewUserService(masterConnectionDB, slaveConnectionDB, logger)
	application := application.NewApplication(userService, logger)

	if err := webserver.NewServer(application).Run(); err != nil {
		logger.Logger.Error(err.Error())
	}
}

