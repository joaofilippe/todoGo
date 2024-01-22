package dto

import (
	"encoding/json"

	"github.com/joaofilippe/todoGo/application/models/users"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Birthdate string `json:"birthdate"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// FromDomain converts a user domain to a user dto
func FromDomain(u users.User) *User {
	return &User{
		ID:        u.ID.String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Birthdate: u.BirthDate.String(),
		Active:    u.Active,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}

// FromJSON converts a json to a user dto
func (u *User) ToJSON() []byte {
	user , _ := json.Marshal(u)
	return user
}
