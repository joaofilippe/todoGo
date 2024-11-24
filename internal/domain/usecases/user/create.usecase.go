package userusecases

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/domain/entities/user"
	"github.com/joaofilippe/todoGo/internal/domain/irepositories"
)

// CreateUseCase is a use case for creating a new user.
type CreateUseCase struct {
	repository irepositories.IUserRepo
}

// NewCreateUserUseCase creates a new instance of CreateUseCase with the given user repository.
func NewCreateUserUseCase(userRepository irepositories.IUserRepo) CreateUseCase {
	return CreateUseCase{
		repository: userRepository,
	}
}

// Execute creates a new user and returns the user's UUID and an error if any.
func (c *CreateUseCase) Execute(newUser user.User) (uuid.UUID, error) {
	return c.repository.CreateNewUser(newUser)
}

