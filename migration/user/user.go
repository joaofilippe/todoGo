package migration

import "github.com/joaofilippe/todoGo/adapters/database/postgres"

// CreateUsersTable creates the users table
func CreateUsersTable(conn *postgres.Connection) error {
	_, err := conn.Connection.Exec(CreateUserTableQuery)
	return err
}