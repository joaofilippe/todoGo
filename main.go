package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"

	"github.com/joaofilippe/todoGo/internal/adapters/database/postgres"
	"github.com/joaofilippe/todoGo/internal/adapters/web"
	"github.com/joaofilippe/todoGo/internal/application"
	"github.com/joaofilippe/todoGo/internal/application/services/users"
	userMigratons "github.com/joaofilippe/todoGo/internal/migrations/users"
	"github.com/joaofilippe/todoGo/pkg/enum"
	"github.com/joaofilippe/todoGo/pkg/logger"
)

func main() {
	env := loadEnv()

	logger := logger.NewLogger(logger.LogOptions{
		Environment: env.ToString(),
	})

	logger.Infof("logando")

	masterConnectionDB := getDbConnection("master", logger)
	slaveConnectionDB := getDbConnection("slave", logger)

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

func loadEnv() enum.Environment {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("can't load .env file")
	}

	return enum.ParseToEnviroment(os.Getenv("ENV"))
}

func getDbConnection(c string, log *logger.Logger) *postgres.Connection {
	yamlFile, err := os.ReadFile(fmt.Sprintf("./config/%s.yaml", c))
	if err != nil {
		log.Fatalf(fmt.Errorf("can't load %s.yaml file", c))
	}

	configDB := new(postgres.Config)

	err = yaml.Unmarshal(yamlFile, configDB)
	if err != nil {
		log.Fatalf(fmt.Errorf("can't unmarshal %s.yaml file", c))
	}

	conn := postgres.NewConnection(configDB)


	if err := conn.Connection.Ping(); err != nil {
		log.Fatalf(fmt.Errorf("can't connect to %s database", c))
	} else {
		log.Infof(fmt.Sprintf("connected to %s database", c))
	}

	return conn
}
