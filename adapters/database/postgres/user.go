package postgres

// Database is a struct that defines the user database
type Database struct {
	MasterConnection *Connection
	SlaveConnection  *Connection
}

// NewDatabase is a function that creates a new user database
func NewDatabase(master, slave *Connection) *Database {
	return &Database{
		MasterConnection: master,
		SlaveConnection:  slave,
	}
}

// CreateUser is a function that creates a user 
func (d *Database) CreateUser()(int64, error){
	

	return 0, nil
}