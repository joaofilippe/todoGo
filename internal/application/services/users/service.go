package users

import (
	"github.com/google/uuid"

	"github.com/joaofilippe/todoGo/internal/adapters/data/postgres"
	consts "github.com/joaofilippe/todoGo/internal/application/consts"
	userEntity "github.com/joaofilippe/todoGo/internal/domain/entities/user"
	userRepo "github.com/joaofilippe/todoGo/internal/application/repository/user"
	"github.com/joaofilippe/todoGo/pkg/logger"
)

// Service is a struct for the user service
type Service struct {
	UserRepository userRepo.IRepository
	UserService    IService
 	Logger         *logger.Logger
}

// NewUserService creates a new user service
func NewUserService(
	writer *postgres.Connection,
	reader *postgres.Connection,
 	logger *logger.Logger,
) *Service {
	return &Service{
		UserRepository: userRepo.NewUserRepository(writer, reader),
 		Logger:         logger,
	}
}

// CreateUser is a usecase to create a new user and returns the id of the new user
func (s *Service) Create(newUser userEntity.NewUser) (uuid.UUID, error) {
	newUser.ID = uuid.New()
	valid, err := newUser.Validate()
	if err != nil {
		return uuid.UUID{}, err
	}

	if !valid {
		return uuid.UUID{}, consts.ErrInvalidNewUser
	}

	exists, err := s.UserRepository.CheckUserExists(newUser.Username, newUser.Email)
	if err != nil {
		return uuid.UUID{}, err
	}

	if exists {
		return uuid.UUID{}, consts.ErrUserAlreadyExists
	}

	err = s.UserRepository.CreateNewUser(newUser)
	if err != nil {
		return uuid.UUID{}, err
	}

	return newUser.ID, nil
}

// Login is a usecase to login a user and returns a token
func (s *Service) Login(login userEntity.Login) (string, error) {
	return "", nil
}

// GetUserByID returns a user by id
func (s *Service) GetUserByID(id uuid.UUID) (userEntity.User, error) {
	return s.UserRepository.GetUserByID(id)
}

// GetUserByUsername returns a user by username
func (s *Service) GetUserByUsername(username string) (userEntity.User, error) {
	return s.UserRepository.GetUserByUsername(username)
}

// GetUserByEmail returns a user by email
func (s *Service) GetUserByEmail(email string) (userEntity.User, error) {
	return s.UserRepository.GetUserByEmail(email)
}
