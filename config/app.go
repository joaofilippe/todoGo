package config

import (
	"errors"
	"os"
	"strings"

	"github.com/joaofilippe/todoGo/pkg/enum"
	"github.com/joaofilippe/todoGo/pkg/logger"
	"github.com/joho/godotenv"
)

type App struct {
	Env        enum.Environment
	ConfigPath string
	SecretKey  string
	Port       string
}

func NewApp(log logger.Logger) *App {
	var app App
	envKey := os.Getenv("ENV")

	if strings.EqualFold(envKey, "") {
		app.Env = enum.Environment(1)
	} else {
		app.Env = enum.ParseToEnviroment(os.Getenv("ENV"))
	}
	
	app.ConfigPath = "../config"
	err := godotenv.Load(app.ConfigPath + "/.env")
	if err != nil {
		log.Fatalf(errors.New("can't load .env file :( "))
	}

	app.SecretKey = os.Getenv("SECRET_KEY")
	app.Port = os.Getenv("PORT")

	return &app
}
