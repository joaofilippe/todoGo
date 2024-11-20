package di_data

import (
	"github.com/joaofilippe/todoGo/config"
	"github.com/joaofilippe/todoGo/internal/adapters/data/database"
	di_connection "github.com/joaofilippe/todoGo/internal/di/connections"

	"github.com/joaofilippe/todoGo/pkg/logger"
)

func GetUserDatabaseWriter(
	logger *logger.Logger,
	appConfig *config.App,
) *database.UserDatabaseWriter {
	masterConn := di_connection.GetConnection(logger, appConfig, "master")
	return &database.UserDatabaseWriter{
		Conn: masterConn,
	}
}
func GetUserDatabaseReader(
	logger *logger.Logger,
	appConfig *config.App,
) *database.UserDatabaseReader {
	slaveConn := di_connection.GetConnection(logger, appConfig, "slave")
	return &database.UserDatabaseReader{
		Conn: slaveConn,
	}
}
