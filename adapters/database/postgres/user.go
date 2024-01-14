package postgres

import "github.com/joaofilippe/todoGo/adapters/database/postgres/query"

// Database is a struct that defines the user database
type Database struct {
	MasterConnection *Connection
	SlaveConnection  *Connection
}

// NewDatabase is a function that creates a new user database
func NewDatabase(master, slave *Connection) *Database {
	return &Database{
		MasterConnection: master,
		SlaveConnection:  slave,
	}
}

// CreateUserTable is a function that creates the user table
func (d *Database) CreateUserTable() error {
	_, err := d.MasterConnection.Connection.Exec(query.CreateUserTableQuery)
	return err
}

// CreateUser is a function that creates a user 
func (d *Database) CreateUser()(int64, error){
	return 0, nil
}