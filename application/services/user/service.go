package user

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/joaofilippe/todoGo/adapters/database/postgres"
	consts "github.com/joaofilippe/todoGo/application/consts"
	userModels "github.com/joaofilippe/todoGo/application/models/user"
	userRepo "github.com/joaofilippe/todoGo/application/repository/user"
	"github.com/joaofilippe/todoGo/common/logger"
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
func (s *Service) CreateUser(newUser *userModels.NewUser) (int, error) {
	userDB, err := s.UserRepository.GetUserByEmail(newUser.Email)
	if err != nil {
		return 0, err
	}

	if userDB.ID != uuid.Nil {
		return 0, errors.New(consts.ErrUserAlreadyExists)
	}

	userID := uuid.New()

	user := &userModels.User{
		ID:        userID,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  newUser.Password,
		Username:  newUser.Username,
		CreatedAt: time.Now(),
		Active:    true,
	}

	return s.UserRepository.CreateUser(user)
}

// Login is a usecase to login a user and returns a token
func (s *Service) Login(login userModels.Login) (string, error) {
	if err := s.Utils.validateLogin(&login); err != nil {
		return "", err
	}

	if login.Email != "" {
		user, err := s.UserRepository.GetUserByEmail(login.Email)
		if err != nil {
			return "", err
		}

		if user.ID == uuid.Nil {
			return "", errors.New(consts.ErrUserDoesNotExist)
		}

		if user.Password != login.Password {
			return "", errors.New(consts.ErrInvalidPassword)
		}

		return s.Utils.generateToken(user)
	}

	user, err := s.UserRepository.GetUserByUsername(login.Username)
	if err != nil {
		return "", err
	}

	if user.ID == uuid.Nil {
		return "", errors.New(consts.ErrUserDoesNotExist)
	}

	if user.Password != login.Password {
		return "", errors.New(consts.ErrInvalidPassword)
	}

	return s.Utils.generateToken(user)
}
