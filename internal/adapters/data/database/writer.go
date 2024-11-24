package database

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/todoGo/internal/adapters/data/database/dto"
	"github.com/joaofilippe/todoGo/internal/adapters/data/database/queries"
	userEntity "github.com/joaofilippe/todoGo/internal/domain/entities/user"
)



// CreateUser is a function that creates a user
func (w *UserDatabaseWriter) CreateNewUser(newUser userEntity.User) (uuid.UUID, error) {
	tx := w.Conn.GetMaster().MustBegin()
	newUserDB := dto.NewUserFromDomain(newUser)

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