package users_migrations

import (
	"github.com/joaofilippe/todoGo/internal/infra/database"
)

// CreateUsersTable creates the users table
func CreateUsersTable(conn *database.Connection) error {
	_, err := conn.GetMaster().Exec(CreateUserTableQuery)
	return err
}

// DropUserTable deletes the users table
func DropUserTable(conn *database.Connection) error {
	_, err := conn.GetMaster().Exec(DropUserTableQuery)
	return err
}
