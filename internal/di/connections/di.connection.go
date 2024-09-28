package di_connection

import (
	"github.com/joaofilippe/todoGo/config"
	"github.com/joaofilippe/todoGo/internal/infra/database"
	"github.com/joaofilippe/todoGo/pkg/enum"
	"github.com/joaofilippe/todoGo/pkg/logger"
	"sync"
)

var (
	masterConnectionDB *infraDatabase.Connection
	slaveConnectionDB  *infraDatabase.Connection
	mu                 sync.Mutex
)

func GetConnection(
	log *logger.Logger,
	appConfig *config.App,
	c string,
) *infraDatabase.Connection {
	mu.Lock()
	defer mu.Unlock()

	if appConfig.Env != enum.Development {
		return &infraDatabase.Connection{}
	}

	switch c {
	case "master":
		if masterConnectionDB == nil {
			masterConnectionDB = infraDatabase.GetConfigFromYaml(log, appConfig, c)
		}
		return masterConnectionDB
	case "slave":
		if slaveConnectionDB == nil {
			slaveConnectionDB = infraDatabase.GetConfigFromYaml(log, appConfig, c)
		}
		return slaveConnectionDB
	default:
		return &infraDatabase.Connection{}
	}
}
