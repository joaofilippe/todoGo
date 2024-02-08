package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"

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
	yamlFile, err := os.ReadFile(fmt.Sprintf("%s/%s.yaml", appConfig.ConfigPath, c))
	if err != nil {
		log.Fatalf(fmt.Errorf("can't load %s.yaml file", c))
	}

	configDB := new(postgres.Config)

	err = yaml.Unmarshal(yamlFile, configDB)
	if err != nil {
		log.Fatalf(fmt.Errorf("can't unmarshal %s.yaml file", c))
	}

	if appConfig.Env != enum.Environment(1) {
		configDB.Host = "host.docker.internal"
	}

	conn := postgres.NewConnection(configDB)

	if err := conn.Connection.Ping(); err != nil {
		fmt.Println(err)
		log.Fatalf(fmt.Errorf("can't connect to %s database", c))
	} else {
		log.Infof(fmt.Sprintf("connected to %s database", c))
	}

	return conn
}
