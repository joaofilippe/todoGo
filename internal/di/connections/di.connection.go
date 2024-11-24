package di_connection

import (
	"sync"

	"github.com/joaofilippe/todoGo/config"
	"github.com/joaofilippe/todoGo/internal/infra/database"
	"github.com/joaofilippe/todoGo/pkg/enum"
	"github.com/joaofilippe/todoGo/pkg/logger"
)

var mu sync.Mutex

// GetConnection returns a database connection based on the application configuration.
func GetConnection(
	log *logger.Logger,
	appConfig *config.App,
) *database.Connection {
	mu.Lock()
	defer mu.Unlock()

	if appConfig.Env() != enum.Development {
		return &database.Connection{}
	}

	return database.New(appConfig)

}
