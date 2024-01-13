package application

import (
	"github.com/joaofilippe/todoGo/application/services/user"
	"github.com/joaofilippe/todoGo/common/logger"
)

// Application represents the application
type Application struct {
	Logger      *logger.Logger
	UserService *user.UserService
}

func NewApplication(userService *user.UserService, logger *logger.Logger) *Application {
	return &Application{
		UserService: userService,
	}
}
