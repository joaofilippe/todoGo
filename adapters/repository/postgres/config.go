package postgres

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"	
)


type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	Dsn      string
} 

type Connection struct {
	Config *Config
	Connection *gorm.DB
}


func (c *Config) getDsn() string {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
						c.Host, 
						c.Port, 
						c.User, 
						c.Password,
						 c.DBName)

	return dsn
}


func NewConnection(config *Config) *Connection {
	var connection *Connection

	config.Dsn = config.getDsn()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		return connection
	}

	connection = &Connection{
		Config: config,
		Connection: db,
	}
	
	return connection
}