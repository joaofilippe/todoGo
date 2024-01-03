package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	common "github.com/joaofilippe/todoGo/common/enum"
	"github.com/joaofilippe/todoGo/common/logger"
)

func main() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("can't load .env file")
	}

	env := common.ParseToEnviroment(os.Getenv("ENV"))
	fmt.Println(env)

	logger := logger.NewLogger(logger.LogOptions{})

	logger.InfoF("logando")
}
