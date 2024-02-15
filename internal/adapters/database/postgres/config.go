package postgres

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //necessary to config
	"gopkg.in/yaml.v3"

	"github.com/joaofilippe/todoGo/config"
	"github.com/joaofilippe/todoGo/pkg/enum"
	"github.com/joaofilippe/todoGo/pkg/logger"
)

type Config struct {
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     string `yaml:"port" env:"DB_PORT"`
	User     string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	DBName   string `yaml:"dbname" env:"DB_NAME"`
	Dsn      string
}

type Connection struct {
	Config     *Config
	Connection *sqlx.DB
}

func (c *Config) getDsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
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

func GetConnection(log *logger.Logger, appConfig *config.App, c string) *Connection {
	if appConfig.Env == enum.Development {
		return GetConfigFromYaml(log, appConfig, c)
	}

	return GetConfigFromEnv()
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

func GetConfigFromEnv() *Connection {
	config := new(Config)

	structE := reflect.ValueOf(config).Elem()
	structT := reflect.TypeOf(*config)

	for i := 0; i < structE.NumField(); i++ {
		fieldT := structT.Field(i)
		fieldName := fieldT.Name
		tagValue := fieldT.Tag.Get("env")
		field := structE.FieldByName(fieldName)

		rs := reflect.ValueOf(os.Getenv(tagValue))
		field.Set(rs)
	}

	return &Connection{}
}
