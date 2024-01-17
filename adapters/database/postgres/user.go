package postgres

import (
	dto "github.com/joaofilippe/todoGo/adapters/database/postgres/dto"
	"github.com/joaofilippe/todoGo/adapters/database/postgres/queries"
	usersModels "github.com/joaofilippe/todoGo/application/models/users"
)

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

// CreateUser is a function that creates a user
func (d *Database) CreateUser(n usersModels.NewUser) (error) {
	newUser := dto.NewUserFromDomain(n)

	_, err := d.MasterConnection.Connection.Exec(queries.InsertNewUserQuery,
		newUser.ID,
		newUser.FirstName,
		newUser.LastName,
		newUser.Username,
		newUser.Email,
		newUser.Password,
		newUser.BirthDate,
		true,
	)
	if err != nil {
		return err
	}

	return err
}
