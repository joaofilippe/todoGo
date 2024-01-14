package application

import (
	"github.com/joaofilippe/todoGo/application/services/user"
	"github.com/joaofilippe/todoGo/common/logger"
)

// Application represents the application
type Application struct {
	Logger      *logger.Logger
	UserService *user.Service
}

func NewApplication(userService *user.Service, logger *logger.Logger) *Application {
	return &Application{
		UserService: userService,
	}
}
