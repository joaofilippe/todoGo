package di_connection

import (
	"sync"

	"github.com/joaofilippe/todoGo/config"
	"github.com/joaofilippe/todoGo/internal/infra/database"
	"github.com/joaofilippe/todoGo/pkg/enum"
	"github.com/joaofilippe/todoGo/pkg/logger"
)

var (
	masterConnectionDB *database.Connection
	slaveConnectionDB  *database.Connection
	mu                 sync.Mutex
)

func GetConnection(
	log *logger.Logger,
	appConfig *config.App,
	c string,
) *database.Connection {
	mu.Lock()
	defer mu.Unlock()

	if appConfig.Env() != enum.Development {
		return &database.Connection{}
	}

	switch c {
	case "master":
		if masterConnectionDB == nil {
			masterConnectionDB = database.GetConfigFromYaml(log, appConfig, c)
		}
		return masterConnectionDB
	case "slave":
		if slaveConnectionDB == nil {
			slaveConnectionDB = database.GetConfigFromYaml(log, appConfig, c)
		}
		return slaveConnectionDB
	default:
		return &database.Connection{}
	}
}
