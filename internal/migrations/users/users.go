package migrations

import "github.com/joaofilippe/todoGo/internal/adapters/database/postgres"

// CreateUsersTable creates the users table
func CreateUsersTable(conn *postgres.Connection) error {
	_, err := conn.Connection.Exec(CreateUserTableQuery)
	return err
}

// DropUserTable deletes the users table
func DropUserTable(conn *postgres.Connection) error {
	_, err := conn.Connection.Exec(DropUserTableQuery)
	return err
}