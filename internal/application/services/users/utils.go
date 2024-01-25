package users

import (
	"os"

	jwt "github.com/golang-jwt/jwt"

	consts "github.com/joaofilippe/todoGo/internal/application/consts"
	usersModels "github.com/joaofilippe/todoGo/internal/application/models/users"
)

type UserUtils struct{}

func (u *UserUtils) validateNewUser(user usersModels.NewUser) error {

	return nil
}

func (u *UserUtils) validateLogin(user usersModels.Login) error {
	if user.Email == "" && user.Username == "" {
		return consts.ErrInvalidLogin
	}

	return nil
}

func (s *UserUtils) generateToken(user usersModels.User) (string, error) {
	secretString := os.Getenv("SECRET_KEY")
	if secretString == "" {
		return "", consts.ErrSecretKey
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"todoIDs":  user.TodoIDs,
	})

	return jwtToken.SignedString([]byte(secretString))
}
