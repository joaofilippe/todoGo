package user

import (
	"errors"
	"time"

	"github.com/google/uuid"

	consts "github.com/joaofilippe/todoGo/application/consts"
	userModels "github.com/joaofilippe/todoGo/application/domain/models/user"
	userRepo "github.com/joaofilippe/todoGo/application/repository/user"
)

type UserService struct {
	UserRepository userRepo.Repository
	UserService    IUserService
	Utils          IUserUtils
}

// CreateUser is a usecase to create a new user and returns the id of the new user
func (s *UserService) CreateUser(newUser *userModels.NewUser) (int, error) {
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

func (s *UserService) Login(login userModels.Login) (string, error) {
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
