package userusecases

import (
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joaofilippe/todoGo/internal/application/consts"
	"github.com/joaofilippe/todoGo/internal/domain/entities/user"
	"github.com/joaofilippe/todoGo/internal/domain/irepositories"
)

type LoginUsecase struct {
	repository irepositories.IUserRepo
}

func NewLoginUseCase(userRepository irepositories.IUserRepo) LoginUsecase {
	return LoginUsecase{
		repository: userRepository,
	}
}

func (l *LoginUsecase) Execute(login user.Login) (string, error) {
	user, err := l.repository.GetUserByEmail(login.Email)
	if err != nil {
		return "", err
	}

	if user.Password != login.Password{
		return "", nil
	}


	return l.generateToken(user)
}

func (l *LoginUsecase) generateToken(user user.User) (string, error) {
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
