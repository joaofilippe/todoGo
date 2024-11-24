package userusecases

import (
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joaofilippe/todoGo/internal/application/consts"
	"github.com/joaofilippe/todoGo/internal/domain/entities"
	"github.com/joaofilippe/todoGo/internal/domain/irepositories"
)

// LoginUsecase handles the user login process.
type LoginUsecase struct {
	repository irepositories.IUserRepo
}

// NewLoginUseCase creates a new instance of LoginUsecase.
func NewLoginUseCase(userRepository irepositories.IUserRepo) LoginUsecase {
	return LoginUsecase{
		repository: userRepository,
	}
}

// Execute handles the login process for a user and returns a JWT token if successful.
func (l *LoginUsecase) Execute(login entities.Login) (string, error) {
	user, err := l.repository.GetUserByEmail(login.Email)
	if err != nil {
		return "", err
	}

	if user.Password != login.Password {
		return "", nil
	}

	return l.generateToken(user)
}

func (l *LoginUsecase) generateToken(user entities.User) (string, error) {
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
