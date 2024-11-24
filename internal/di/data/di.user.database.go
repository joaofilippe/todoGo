package datadi

import (
	"github.com/joaofilippe/todoGo/config"
	userdatabase "github.com/joaofilippe/todoGo/internal/adapters/data/database/user"
	di_connection "github.com/joaofilippe/todoGo/internal/di/connections"

	"github.com/joaofilippe/todoGo/pkg/logger"
)

// GetUserDatabaseWriter returns a new UserDatabaseWriter with a master connection.
func GetUserDatabaseWriter(
	logger *logger.Logger,
	appConfig *config.App,
) *userdatabase.Writer {
	masterConn := di_connection.GetConnection(logger, appConfig)
	return &userdatabase.Writer{
		Conn: masterConn,
	}
}

// GetUserDatabaseReader returns a new UserDatabaseReader with a slave connection.
func GetUserDatabaseReader(
	logger *logger.Logger,
	appConfig *config.App,
) *userdatabase.Reader {
	slaveConn := di_connection.GetConnection(logger, appConfig)
	return &userdatabase.Reader{
		Conn: slaveConn,
	}
}
