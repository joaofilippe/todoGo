package user

import userModels "github.com/joaofilippe/todoGo/application/domain/models/user"

type UserService struct {
	UserRepository IUserRepository
	UserService	IUserService
}

// CreateUser is a usecase to create a new user and returns the id of the new user
func (s *UserService) CreateUser(newUser *userModels.NewUser) (int, error) {	
	user := &userModels.User{}
	return s.UserRepository.Create(user)
}