package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"

	"github.com/joaofilippe/todoGo/pkg/enum"
	"github.com/joaofilippe/todoGo/pkg/logger"
)

var app *App

// App represents the application configuration and state.
type App struct {
	env        enum.Environment
	Logger     *logger.Logger
	configPath string
	secretKey  string
	port       string
	masterDsn  string
	slaveDsn   string
}

type dbConfig struct {
	host     string `yaml:"host" env:"DB_HOST"`
	port     string `yaml:"port" env:"DB_PORT"`
	user     string `yaml:"user" env:"DB_USER"`
	password string `yaml:"password" env:"DB_PASSWORD"`
	dbName   string `yaml:"dbname" env:"DB_NAME"`
}

// New creates a new application instance with the provided logger.
func New(log *logger.Logger) *App {
	if app != nil {
		return app
	}

	app = &App{
		Logger: log,
	}

	envKey := os.Getenv("ENV")

	app.env = enum.ParseToEnviroment(envKey)

	app.configPath = "../config"

	err := godotenv.Load(app.configPath + "/.env")
	if err != nil {
		log.Fatalf(errors.New("can't load .env file :( "))
	}

	app.secretKey = os.Getenv("SECRET_KEY")
	app.port = os.Getenv("PORT")
	app.setDsn()

	return app
}

func (a *App) setDsn() {
	masterDbConfig := &dbConfig{}
	slaveDbConfig := &dbConfig{}
	if a.env == enum.Development {
		a.getConfigFromYaml(masterDbConfig, "master")
		a.getConfigFromYaml(slaveDbConfig, "slave")
	}

	a.masterDsn = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		masterDbConfig.host,
		masterDbConfig.port,
		masterDbConfig.user,
		masterDbConfig.password,
		masterDbConfig.dbName,
	)
	a.slaveDsn = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		slaveDbConfig.host,
		slaveDbConfig.port,
		slaveDbConfig.user,
		slaveDbConfig.password,
		slaveDbConfig.dbName,
	)
}

func (a *App) getConfigFromYaml(dbConfig *dbConfig, dbRelation string) {
	yamlFile, err := os.ReadFile(fmt.Sprintf("%s/%s.yaml", a.configPath, dbRelation))
	if err != nil {
		a.Logger.Fatalf(fmt.Errorf("can't load %s.yaml file", dbRelation))
	}

	err = yaml.Unmarshal(yamlFile, dbConfig)
	if err != nil {
		a.Logger.Fatalf(fmt.Errorf("can't unmarshal %s.yaml file", dbRelation))
	}
}

func (a *App) MasterDsn() string {
	return a.masterDsn
}

func (a *App) SlaveDsn() string {
	return a.slaveDsn
}

func (a *App) Env() enum.Environment {
	return a.env
}
