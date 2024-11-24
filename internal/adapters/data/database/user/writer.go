package userdatabase

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/adapters/data/database/queries"
	"github.com/joaofilippe/todoGo/internal/domain/entities"
	"github.com/joaofilippe/todoGo/internal/infra/database"
)

// Writer is a struct that holds the database connection
type Writer struct {
	Conn *database.Connection
}

// CreateNewUser is a function that creates a user
func (w *Writer) CreateNewUser(newUser entities.User) (uuid.UUID, error) {
	tx := w.Conn.GetMaster().MustBegin()
	newUserDB := NewUserFromDomain(newUser)

	_, err :=
		tx.Exec(
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
		tx.Rollback()
		return uuid.UUID{}, err
	}

	err = tx.Commit()
	return newUser.ID, err
}
