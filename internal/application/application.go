package application

import (
	"github.com/joaofilippe/todoGo/internal/infra/user_service"
	"github.com/joaofilippe/todoGo/pkg/logger"
)

// Application represents the application
type Application struct {
	Logger      *logger.Logger
	UserService *user_service.UserService
}

// NewApplication creates a new application
func NewApplication(userService *user_service.UserService, logger *logger.Logger) *Application {
	return &Application{
		UserService: userService,
	}
}
