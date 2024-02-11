package main

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/joaofilippe/todoGo/config"
	"github.com/joaofilippe/todoGo/internal/adapters/database/postgres"
	"github.com/joaofilippe/todoGo/internal/adapters/web"
	"github.com/joaofilippe/todoGo/internal/application"
	"github.com/joaofilippe/todoGo/internal/application/services/users"
	userMigratons "github.com/joaofilippe/todoGo/internal/migrations/users"
	"github.com/joaofilippe/todoGo/pkg/enum"
	"github.com/joaofilippe/todoGo/pkg/logger"
)

func main() {
	logger := logger.NewLogger()
	appConfig := config.NewApp(*logger)

	masterConnectionDB := getDbConnection(logger, appConfig, "master")
	slaveConnectionDB := getDbConnection(logger, appConfig, "slave")

	if err := userMigratons.CreateUsersTable(masterConnectionDB); err != nil {
		logger.Logger.Error(err.Error())
	}

	userUtils := &users.UserUtils{}
	userService := users.NewUserService(masterConnectionDB, slaveConnectionDB, userUtils, logger)
	application := application.NewApplication(userService, logger)

	if err := webserver.NewServer(application).Run(); err != nil {
		logger.Logger.Error(err.Error())
	}

	fmt.Println(masterConnectionDB, slaveConnectionDB)
}

func getDbConnection(log *logger.Logger, appConfig *config.App, c string) *postgres.Connection {
	if appConfig.Env == enum.Development {
		return postgres.GetConfigFromYaml(log, appConfig,c)
	}

	return &postgres.Connection{}
}
