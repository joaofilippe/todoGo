package application

import (
	"github.com/joaofilippe/todoGo/internal/infra/services"
 	"github.com/joaofilippe/todoGo/pkg/logger"
)

// Application represents the application
type Application struct {
	Logger      *logger.Logger
	UserService *services.UserService
}

// NewApplication creates a new application
func NewApplication(userService *services.UserService, logger *logger.Logger) *Application {
	return &Application{
		UserService: userService,
	}
}
