package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	dto "github.com/joaofilippe/todoGo/internal/adapters/database/postgres/dto"
	models "github.com/joaofilippe/todoGo/internal/adapters/database/postgres/models"
	"github.com/joaofilippe/todoGo/internal/adapters/database/postgres/queries"
	userEntity "github.com/joaofilippe/todoGo/internal/application/entities/user"
)

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
func (d *Database) CreateNewUser(n userEntity.NewUser) (error) {
	newUser := dto.NewUserFromDomain(n)

	_, err := d.MasterConnection.Connection.Exec(queries.InsertNewUserQuery,
		newUser.ID,
		newUser.FirstName,
		newUser.LastName,
		newUser.Username,
		newUser.Email,
		newUser.Password,
		newUser.BirthDate,
		true,
	)
	if err != nil {
		return err
	}

	return err
}

// GetUserByID is a function that gets a user by username
func (d *Database) GetUserByID(id uuid.UUID) (userEntity.User, error) {
	var userDB models.UserDB

	err := d.SlaveConnection.Connection.Get(&userDB, queries.SelectUserQueryByID, id.String())
	if err != nil && err != sql.ErrNoRows{
		return userEntity.User{}, err
	}

	return dto.UserFromDB(userDB).ToDomain(), nil
}

// GetUserByEmail is a function that gets a user by username
func (d *Database) GetUserByEmail(email string) (userEntity.User, error) {
	var userDB models.UserDB

	err := d.SlaveConnection.Connection.Get(&userDB, queries.SelectUserQueryByEmail, email)
	if err == sql.ErrNoRows {
		return userEntity.User{}, nil
	}

	if err != nil{
		return userEntity.User{}, err
	}

	

	return dto.UserFromDB(userDB).ToDomain(), nil
}

// GetUserByUsername is a function that gets a user by username
func (d *Database) GetUserByUsername(username string) (userEntity.User, error) {
	var userDB models.UserDB

	err := d.SlaveConnection.Connection.Get(&userDB, queries.SelectUserQueryByUsername, username)
	if err != nil && err != sql.ErrNoRows{
		return userEntity.User{}, err
	}

	return dto.UserFromDB(userDB).ToDomain(), nil
}

// CheckUserExists is a function that checks if a user exists
func (d *Database) CheckUserExists(username, email string) (bool, error) {
	var exists bool

	err := d.SlaveConnection.Connection.Get(&exists, queries.CheckUserExists, username, email)
	if err != nil {
		return false, err
	}

	return exists, nil
}