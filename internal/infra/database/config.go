package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joaofilippe/todoGo/config"
	_ "github.com/lib/pq" //necessary to config
)

var connection *Connection

type Connection struct {
	master *sqlx.DB
	slave  *sqlx.DB
}

// New creates a new database connection using the provided app configuration.
func New(app *config.App) *Connection {
	if connection != nil {
		return connection
	}

	masterDB, err := sqlx.Open("postgres", app.MasterDsn())
	if err != nil {
		log.Fatalf(err.Error())
		return connection
	}

	slaveDB, err := sqlx.Open("postgres", app.SlaveDsn())
	if err != nil {
		log.Fatalf(err.Error())
		return connection
	}

	connection.master = masterDB
	connection.slave = slaveDB

	return connection
}

// GetMaster returns the master database connection.
func (c *Connection) GetMaster() *sqlx.DB {
	return c.master
}

// GetSlave returns the slave database connection.
func (c *Connection) GetSlave() *sqlx.DB {
	return c.slave
}
