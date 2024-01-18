package users

import (
	"github.com/google/uuid"

	"github.com/joaofilippe/todoGo/adapters/database/postgres"
	consts "github.com/joaofilippe/todoGo/application/consts"
	usersModels "github.com/joaofilippe/todoGo/application/models/users"
	userRepo "github.com/joaofilippe/todoGo/application/repository/user"
	"github.com/joaofilippe/todoGo/pkg/logger"
)

// Service is a struct for the user service
type Service struct {
	UserRepository userRepo.IRepository
	UserService    IUserService
	Utils          IUserUtils
	Logger         *logger.Logger
}

// NewUserService creates a new user service
func NewUserService(
	writer *postgres.Connection,
	reader *postgres.Connection,
	utils IUserUtils,
	logger *logger.Logger,
) *Service {
	return &Service{
		UserRepository: userRepo.NewUserRepository(writer, reader),
		Utils:          utils,
		Logger:         logger,
	}
}

// CreateUser is a usecase to create a new user and returns the id of the new user
func (s *Service) CreateUser(newUser usersModels.NewUser) (uuid.UUID, error) {
	newUser.ID = uuid.New()
	valid, err := newUser.Validate()
	if err != nil {
		return uuid.UUID{}, err
	}

	if !valid {
		return uuid.UUID{}, consts.ErrInvalidNewUser
	}

	userDB, err := s.UserRepository.GetUserByEmail(newUser.Email)
	if err != nil {
		return uuid.UUID{}, err
	}

	if userDB.ID != uuid.Nil {
		return uuid.UUID{}, consts.ErrUserAlreadyExists
	}
	
	err = s.UserRepository.CreateNewUser(newUser)
	if err != nil {
		return uuid.UUID{}, err
	}

	return newUser.ID, nil
}

// Login is a usecase to login a user and returns a token
func (s *Service) Login(login usersModels.Login) (string, error) {
	if err := s.Utils.validateLogin(login); err != nil {
		return "", err
	}

	if login.Email != "" {
		user, err := s.UserRepository.GetUserByEmail(login.Email)
		if err != nil {
			return "", err
		}

		if user.ID == uuid.Nil {
			return "", consts.ErrUserDoesNotExist
		}

		if user.Password != login.Password {
			return "", consts.ErrInvalidPassword
		}

		return s.Utils.generateToken(user)
	}

	user, err := s.UserRepository.GetUserByUsername(login.Username)
	if err != nil {
		return "", err
	}

	if user.ID == uuid.Nil {
		return "", consts.ErrUserDoesNotExist
	}

	if user.Password != login.Password {
		return "", consts.ErrInvalidPassword
	}

	return s.Utils.generateToken(user)
}
