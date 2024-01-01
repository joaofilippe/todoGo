package user

import (
	"errors"

	consts "github.com/joaofilippe/todoGo/application/consts"
	userModels "github.com/joaofilippe/todoGo/application/domain/models/user"
)

type UserUtils struct {}

func (u *UserUtils) validateUser(user *userModels.NewUser) error {

	return nil
}

func (u *UserUtils) validateLogin(user *userModels.Login) error {
	if user.Email == "" && user.Username == "" {
	return errors.New(consts.ErrInvalidLogin)
	}

	return nil
}

func (s *UserUtils) generateToken(user *userModels.User) (string, error) {
	return "", nil
}
