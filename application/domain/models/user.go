package models

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

func (u *User) CreateNewUser(
	firstName, 
	lastName, 
	username, 
	email, 
	password string) {

	u.FirstName = firstName
	u.LastName = lastName
	u.Username = username
	u.Email = email
	u.Password = password

	
}