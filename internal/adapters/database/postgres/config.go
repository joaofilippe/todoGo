package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joaofilippe/todoGo/config"
	"github.com/joaofilippe/todoGo/pkg/enum"
	"github.com/joaofilippe/todoGo/pkg/logger"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Dsn      string
}

type Connection struct {
	Config     *Config
	Connection *sqlx.DB
}

func (c *Config) getDsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName)

}

func NewConnection(config *Config) *Connection {
	var connection *Connection

	config.Dsn = config.getDsn()

	db, err := sqlx.Open("postgres", config.Dsn)
	if err != nil {
		log.Fatalf(err.Error())
		return connection
	}

	connection = &Connection{
		Config:     config,
		Connection: db,
	}

	return connection
}

func GetConfigFromYaml(log *logger.Logger, appConfig *config.App, c string) *Connection {
	yamlFile, err := os.ReadFile(fmt.Sprintf("%s/%s.yaml", appConfig.ConfigPath, c))
	if err != nil {
		log.Fatalf(fmt.Errorf("can't load %s.yaml file", c))
	}

	configDB := new(Config)

	err = yaml.Unmarshal(yamlFile, configDB)
	if err != nil {
		log.Fatalf(fmt.Errorf("can't unmarshal %s.yaml file", c))
	}

	if appConfig.Env != enum.Environment(1) {
		configDB.Host = "host.docker.internal"
	}

	conn := NewConnection(configDB)

	if err := conn.Connection.Ping(); err != nil {
		log.Fatalf(fmt.Errorf("can't connect to %s database. Error: %s", c, err.Error()))
	} else {
		log.Fatalf(fmt.Errorf("can't connect to %s database. Error: %s", c, err.Error()))
	}

	return conn
}

func GetConfigFromEnv(log *logger.Logger, appConfig *config.App, c string) *Connection {

	return &Connection{}
}
