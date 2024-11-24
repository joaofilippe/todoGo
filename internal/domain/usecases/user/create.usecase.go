package userusecases

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/domain/entities"
	"github.com/joaofilippe/todoGo/internal/domain/irepositories"
)

// CreateUseCase is a use case for creating a new user.
type CreateUseCase struct {
	repository irepositories.IUserRepo
}

// NewCreateUseCase creates a new instance of CreateUseCase with the given user repository.
func NewCreateUseCase(userRepository irepositories.IUserRepo) CreateUseCase {
	return CreateUseCase{
		repository: userRepository,
	}
}

// Execute creates a new user and returns the user's UUID and an error if any.
func (c *CreateUseCase) Execute(newUser entities.User) (uuid.UUID, error) {
	return c.repository.CreateNewUser(newUser)
}
