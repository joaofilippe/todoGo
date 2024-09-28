package migrations

import (
	infraDatabase "github.com/joaofilippe/todoGo/internal/infra/database"
)

// CreateUsersTable creates the users table
func CreateUsersTable(conn *infraDatabase.Connection) error {
	_, err := conn.Connection.Exec(CreateUserTableQuery)
	return err
}

// DropUserTable deletes the users table
func DropUserTable(conn *infraDatabase.Connection) error {
	_, err := conn.Connection.Exec(DropUserTableQuery)
	return err
}
