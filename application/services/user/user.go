package user

import (
	"errors"

	consts "github.com/joaofilippe/todoGo/application/consts"
	userModels "github.com/joaofilippe/todoGo/application/domain/models/user"
)

func (s *UserService) ValidateUser(user *userModels.NewUser) error {

	return nil
}

func (s *UserService) ValidateLogin(user *userModels.Login) error {
	if user.Email == "" && user.Username == "" {
		return errors.New(consts.ErrInvalidLogin)
	}

	return nil
}

func (s *UserService) GenerateToken(user *userModels.User) (string, error) {
	return "", nil
}
