package userservice

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/domain/entities"
	"github.com/joaofilippe/todoGo/internal/domain/usecases/user"
)

// UserService provides user-related services.
type UserService struct {
	createUsecase userusecases.CreateUseCase
	loginUsecase  userusecases.LoginUsecase
}

// NewUserService creates a new instance of UserService.
func NewUserService(
	create userusecases.CreateUseCase,
	login userusecases.LoginUsecase,
) *UserService {
	return &UserService{}
}

// Create creates a new user and returns the user's UUID.
func (s *UserService) Create(newUser entities.User) (uuid.UUID, error) {
	return s.createUsecase.Execute(newUser)
}

// Login logs in a user and returns a JWT token.
func (s *UserService) Login(login entities.Login) (string, error) {
	return s.loginUsecase.Execute(login)
}

// GetByID retrieves a user by their ID.
func (s *UserService) GetByID(id uuid.UUID) (entities.User, error) {
	return entities.User{}, nil
}

// GetByEmail retrieves a user by their email.
func (s *UserService) GetByEmail(email string) (entities.User, error) {
	return entities.User{}, nil
}

// GetByUsername retrieves a user by their username.
func (s *UserService) GetByUsername(username string) (entities.User, error) {
	return entities.User{}, nil
}
