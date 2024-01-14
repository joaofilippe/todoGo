package postgres

// UserRepository is a struct that defines the user repository
type UserRepository struct {
	Connection *Connection
}

// NewUserRepository is a function that creates a new user repository
func NewUserRepository(connection *Connection) *UserRepository {
	return &UserRepository{Connection: connection}
}

// CreateUserTable is a function that creates the user table
func CreateUserTable() error {
	return nil
}

// CreateUser is a function that creates a user 
func (r *UserRepository) CreateUser()(int64, error){
	return 0, nil
}