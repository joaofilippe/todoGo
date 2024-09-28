package user

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/domain/entities/user"
	userRepo "github.com/joaofilippe/todoGo/internal/domain/repositories"
)

type CreateUseCase struct {
	repository *userRepo.IUserRepo
}

func NewCreateUserUseCase(userRepository *userRepo.IUserRepo) CreateUseCase {
	return CreateUseCase{
		repository: userRepository,
	}
}

func (c *CreateUseCase) Execute(newUser user.NewUser) (uuid.UUID, error) {
	return c.repository.Writer.Create(newUser)
}

