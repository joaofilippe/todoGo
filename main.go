package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"

	"github.com/joaofilippe/todoGo/adapters/repository/postgres"
	common "github.com/joaofilippe/todoGo/common/enum"
	"github.com/joaofilippe/todoGo/common/logger"
)

func main() {
	env := loadEnv()

	logger := logger.NewLogger(logger.LogOptions{
		Environment: env.ToString(),
	})

	logger.InfoF("logando")

	masterConnectionDB := getDbConnection("master")
	slaveConnectionDB := getDbConnection("slave")

	fmt.Println(masterConnectionDB, slaveConnectionDB)
}

func loadEnv() common.Environment{
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("can't load .env file")
	}

	return common.ParseToEnviroment(os.Getenv("ENV"))
}

func getDbConnection(c string) *postgres.Connection {
	yamlFile, err := os.ReadFile(fmt.Sprintf("./config/%s.yaml", c))
	if err != nil {
		log.Fatalf("can't load %s.yaml file", c)
	}

	configDB := new(postgres.Config)

	err = yaml.Unmarshal(yamlFile, configDB)
	if err != nil {
		log.Fatalf("can't unmarshal %s.yaml file", c)
	}

	conn := postgres.NewConnection(configDB)
	defer conn.Connection.Close()
	
	if err := conn.Connection.Ping(); err != nil {
		log.Fatalf("can't connect to %s database", c)
	} else {
		log.Printf("connected to %s database", c)
	}

	return conn
}

