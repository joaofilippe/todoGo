package iuser

// IUser is an interface that defines the user database
type IUser interface {
	CreateUser() (int64, error)
	CreateUserTable() error
}