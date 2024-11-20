package di_repositories

import (
	"github.com/joaofilippe/todoGo/config"
	"github.com/joaofilippe/todoGo/internal/adapters/data"
	di_data "github.com/joaofilippe/todoGo/internal/di/data"
	"github.com/joaofilippe/todoGo/pkg/logger"
)

func GetUserRepository(
	logger *logger.Logger,
	appConfig *config.App,
) *data.UserRepository {
	writer := di_data.GetUserDatabaseWriter(logger, appConfig)
	reader := di_data.GetUserDatabaseReader(logger, appConfig)

	return &data.UserRepository{
		Writer: writer,
		Reader: reader,
	}
}
