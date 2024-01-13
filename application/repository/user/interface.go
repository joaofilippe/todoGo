package user

import (
	userModels "github.com/joaofilippe/todoGo/application/models/user"
)

// Repository represents the user repository
type Repository interface {
	RepositoryReader
	RepositoryWriter
}

// RepositoryWriter represents the user repository writer
type RepositoryWriter interface {
	CreateUser(user *userModels.User) (int, error)
}

// RepositoryReader represents the user repository reader
type RepositoryReader interface {
	GetUserByUsername(username string) (*userModels.User, error)
	GetUserByEmail(email string) (*userModels.User, error)
}
