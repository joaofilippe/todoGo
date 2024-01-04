package postgres

type UserRepository struct {
	Connection *Connection
}

func (r *UserRepository) Create()(int64, error){
	return 0, nil
}