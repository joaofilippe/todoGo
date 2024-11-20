package database

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/infra/database"
	"github.com/joaofilippe/todoGo/internal/adapters/data/database/dto"
	"github.com/joaofilippe/todoGo/internal/adapters/data/database/models"
	"github.com/joaofilippe/todoGo/internal/adapters/data/database/queries"
	userEntity "github.com/joaofilippe/todoGo/internal/domain/entities/user"
)

type IUserWriter interface {
	CreateNewUser(newUser userEntity.User) (uuid.UUID, error)
}

type IUserReader interface {
	GetUserByID(id uuid.UUID) (userEntity.User, error)
	GetUserByEmail(email string) (userEntity.User, error)
	GetUserByUsername(username string) (userEntity.User, error)
}

type UserDatabaseWriter struct {
	Conn *infraDatabase.Connection
}

type UserDatabaseReader struct {
	Conn *infraDatabase.Connection
}

// CreateUser is a function that creates a user
func (w *UserDatabaseWriter) CreateNewUser(newUser userEntity.User) (uuid.UUID, error) {
	newUserDB := dto.NewUserFromDomain(newUser)

	_, err :=
		w.Conn.Connection.Exec(
			queries.InsertNewUserQuery,
			newUserDB.ID,
			newUserDB.FirstName,
			newUserDB.LastName,
			newUserDB.Username,
			newUserDB.Email,
			newUserDB.Password,
			newUserDB.BirthDate,
			true,
		)
	if err != nil {
		return uuid.UUID{}, err
	}

	return newUser.ID, err
}

// GetUserByID is a function that gets a user by username
func (r *UserDatabaseReader) GetUserByID(id uuid.UUID) (userEntity.User, error) {
	var userDB models.UserDB

	err := r.Conn.Connection.Get(&userDB, queries.SelectUserQueryByID, id.String())
	if err != nil && err != sql.ErrNoRows {
		return userEntity.User{}, err
	}

	return dto.UserFromDB(userDB).ToDomain(), nil
}

// GetUserByEmail is a function that gets a user by username
func (r *UserDatabaseReader) GetUserByEmail(email string) (userEntity.User, error) {
	var userDB models.UserDB

	err := r.Conn.Connection.Get(&userDB, queries.SelectUserQueryByEmail, email)
	if err == sql.ErrNoRows {
		return userEntity.User{}, nil
	}

	if err != nil {
		return userEntity.User{}, err
	}

	return dto.UserFromDB(userDB).ToDomain(), nil
}

// GetUserByUsername is a function that gets a user by username
func (r *UserDatabaseReader) GetUserByUsername(username string) (userEntity.User, error) {
	var userDB models.UserDB

	err := r.Conn.Connection.Get(&userDB, queries.SelectUserQueryByUsername, username)
	if err != nil && err != sql.ErrNoRows {
		return userEntity.User{}, err
	}

	return dto.UserFromDB(userDB).ToDomain(), nil
}
