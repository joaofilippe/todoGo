package userdatabase

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	entities "github.com/joaofilippe/todoGo/internal/domain/entities"
	uuid_helper "github.com/joaofilippe/todoGo/pkg/uuid"
)

// UserDTO is the model for the user table
type UserDTO struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	TodoIDs   []uuid.UUID
	BirthDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Active    bool
}

// UserFromDomain is a function that creates a new UserDTO
func UserFromDomain(u entities.User) *UserDTO {
	return &UserDTO{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		TodoIDs:   u.TodoIDs,
		BirthDate: u.BirthDate,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Active:    u.Active,
	}
}

// UserFromDB converts the UserDB to a UserDTO
func UserFromDB(user UserDB) *UserDTO {
	var todoIDs []uuid.UUID
	if user.TodoIDs.Valid {
		todoIDs, _ = uuid_helper.ConvertStringToArray(user.TodoIDs.String[1 : len(user.TodoIDs.String)-1])
	}

	birthDate, _ := time.Parse(time.RFC3339, user.BirthDate.String)
	createdAt, _ := time.Parse(time.RFC3339, user.CreatedAt.String)
	updatedAt, _ := time.Parse(time.RFC3339, user.UpdatedAt.String)

	return &UserDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		TodoIDs:   todoIDs,
		BirthDate: birthDate,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Active:    user.Active,
	}
}

// ToDomain converts the UserDTO to a User domain model
func (u *UserDTO) ToDomain() entities.User {
	return entities.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		TodoIDs:   u.TodoIDs,
		BirthDate: u.BirthDate,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Active:    u.Active,
	}
}

// ToDB converts the UserDTO to a UserDB
func (u *UserDTO) ToDB() UserDB {
	return UserDB{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		TodoIDs:   sql.NullString{},
		BirthDate: sql.NullString{},
		CreatedAt: sql.NullString{},
		UpdatedAt: sql.NullString{},
		Active:    u.Active,
	}
}
