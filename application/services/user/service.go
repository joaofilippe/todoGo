package user

import "todoGo/application/domain/models"

type UserService struct {
	UserRepository IUserRepository
	UserService	IUserService
}

func (s *UserService) Create(user *models.User) (int, error) {
	return s.UserRepository.Create(user)
}