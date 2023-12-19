package user

import "todoGo/application/domain/models"

type UserService struct {
	UserRepository IUserRepository
	UserService	IUserService
}

// CreateUser is a usecase to create a new user and returns the id of the new user
func (s *UserService) CreateUser(user *models.User) (int, error) {
	

	return s.UserRepository.Create(user)
}