package application

import (
	"github.com/joaofilippe/todoGo/internal/application/services/users"
	"github.com/joaofilippe/todoGo/pkg/logger"
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
