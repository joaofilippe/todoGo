package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"

	"github.com/joaofilippe/todoGo/adapters/repository/postgres"
	"github.com/joaofilippe/todoGo/adapters/web"
	"github.com/joaofilippe/todoGo/application"
	"github.com/joaofilippe/todoGo/application/services/user"
	common "github.com/joaofilippe/todoGo/common/enum"
	"github.com/joaofilippe/todoGo/common/logger"
)

func main() {
	env := loadEnv()

	logger := logger.NewLogger(logger.LogOptions{
		Environment: env.ToString(),
	})

	logger.Infof("logando")

	masterConnectionDB := getDbConnection("master", logger)
	slaveConnectionDB := getDbConnection("slave", logger)

	userUtils := &user.UserUtils{}
	userService := user.NewUserService(masterConnectionDB, slaveConnectionDB, userUtils, logger)
	application := application.NewApplication(userService, logger)

	server := webserver.NewServer(application)

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8080"
	}

	if err := http.ListenAndServe(":"+PORT, server.Router); err != nil {
		logger.Logger.Error(err.Error())
	}

	fmt.Println(masterConnectionDB, slaveConnectionDB)
}

func loadEnv() common.Environment {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("can't load .env file")
	}

	return common.ParseToEnviroment(os.Getenv("ENV"))
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
	defer conn.Connection.Close()

	if err := conn.Connection.Ping(); err != nil {
		log.Fatalf(fmt.Errorf("can't connect to %s database", c))
	} else {
		log.Infof(fmt.Sprintf("connected to %s database", c))
	}

	return conn
}
