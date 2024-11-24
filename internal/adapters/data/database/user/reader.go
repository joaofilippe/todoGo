package userdatabase

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/adapters/data/database/queries"
	userEntity "github.com/joaofilippe/todoGo/internal/domain/entities/user"
	"github.com/joaofilippe/todoGo/internal/infra/database"
)


// Reader is a struct that holds the database connection
type Reader struct {
	Conn *database.Connection
}


// GetUserByID is a function that gets a user by username
func (r *Reader) GetUserByID(id uuid.UUID) (userEntity.User, error) {
	tx := r.Conn.GetSlave().MustBegin()
	var userDB UserDB

	err := tx.Get(&userDB, queries.SelectUserQueryByID, id.String())
	if err != nil && err != sql.ErrNoRows {
		return userEntity.User{}, err
	}

	return UserFromDB(userDB).ToDomain(), nil
}

// GetUserByEmail is a function that gets a user by username
func (r *Reader) GetUserByEmail(email string) (userEntity.User, error) {
	tx := r.Conn.GetSlave().MustBegin()
	var userDB UserDB

	err := tx.Get(&userDB, queries.SelectUserQueryByEmail, email)
	if err == sql.ErrNoRows {
		return userEntity.User{}, nil
	}

	if err != nil {
		tx.Rollback()
		return userEntity.User{}, err
	}

	err = tx.Commit()
	return UserFromDB(userDB).ToDomain(), err
}

// GetUserByUsername is a function that gets a user by username
func (r *Reader) GetUserByUsername(username string) (userEntity.User, error) {
	tx := r.Conn.GetSlave().MustBegin()
	var userDB UserDB

	err := tx.Get(&userDB, queries.SelectUserQueryByUsername, username)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return userEntity.User{}, err
	}

	err = tx.Commit()
	return UserFromDB(userDB).ToDomain(), err
}
