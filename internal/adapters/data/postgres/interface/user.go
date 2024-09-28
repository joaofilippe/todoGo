package iuser

// IUserRepo is an interface that defines the user database
type IUserRepo interface {
	CreateUser() (int64, error)
	CreateUserTable() error
}