package users

import (
	"os"

	jwt "github.com/golang-jwt/jwt"

	consts "github.com/joaofilippe/todoGo/internal/application/consts"
	userEntity "github.com/joaofilippe/todoGo/internal/application/entities/user"
)

type UserUtils struct{}

func (u *UserUtils) validateNewUser(user userEntity.NewUser) error {

	return nil
}

func (u *UserUtils) validateLogin(user userEntity.Login) error {
	if user.Email == "" && user.Username == "" {
		return consts.ErrInvalidLogin
	}

	return nil
}

func (s *UserUtils) generateToken(user userEntity.User) (string, error) {
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
