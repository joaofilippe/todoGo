package user

import (
	"errors"
	"os"

	jwt "github.com/golang-jwt/jwt"

	consts "github.com/joaofilippe/todoGo/application/consts"
	userModels "github.com/joaofilippe/todoGo/application/models/user"
)

type UserUtils struct{}

func (u *UserUtils) validateNewUser(user *userModels.NewUser) error {

	return nil
}

func (u *UserUtils) validateLogin(user *userModels.Login) error {
	if user.Email == "" && user.Username == "" {
		return errors.New(consts.ErrInvalidLogin)
	}

	return nil
}

func (s *UserUtils) generateToken(user *userModels.User) (string, error) {
	secretString := os.Getenv("SECRET_KEY")
	if secretString == "" {
		return "", errors.New(consts.ErrSecretKey)
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"todoIDs":  user.TodoIDs,
	})

	return jwtToken.SignedString([]byte(secretString))
}
