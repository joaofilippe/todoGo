package di_data

import (
	"github.com/joaofilippe/todoGo/config"
	"github.com/joaofilippe/todoGo/internal/adapters/data/database"
	di_connection "github.com/joaofilippe/todoGo/internal/di/connections"

	"github.com/joaofilippe/todoGo/pkg/logger"
)

// GetUserDatabaseWriter returns a new UserDatabaseWriter with a master connection.
func GetUserDatabaseWriter(
	logger *logger.Logger,
	appConfig *config.App,
) *database.UserDatabaseWriter {
	masterConn := di_connection.GetConnection(logger, appConfig)
	return &database.UserDatabaseWriter{
		Conn: masterConn,
	}
}
// GetUserDatabaseReader returns a new UserDatabaseReader with a slave connection.
func GetUserDatabaseReader(
	logger *logger.Logger,
	appConfig *config.App,
) *database.UserDatabaseReader {
	slaveConn := di_connection.GetConnection(logger, appConfig)
	return &database.UserDatabaseReader{
		Conn: slaveConn,
	}
}
