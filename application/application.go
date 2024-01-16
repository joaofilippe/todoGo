package application

import (
	"github.com/joaofilippe/todoGo/application/services/users"
	"github.com/joaofilippe/todoGo/common/logger"
)

// Application represents the application
type Application struct {
	Logger      *logger.Logger
	UserService *users.Service
}

// NewApplication creates a new application
func NewApplication(userService *users.Service, logger *logger.Logger) *Application {
	return &Application{
		UserService: userService,
	}
}
