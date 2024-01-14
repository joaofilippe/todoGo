package user

import (
	"net/http"

	webserver "github.com/joaofilippe/todoGo/adapters/web"
)

var (
	// ErrCannotCreateUser is the default cannot create user error
	ErrCannotCreateUser = &webserver.ErrResponse{HTTPStatusCode: 400, StatusText: "Cannot create user."}

	// ErrUserNotFound is the default user not found error
	ErrUserNotFound = &webserver.ErrResponse{HTTPStatusCode: 404, StatusText: "User not found."}
)

var (
	// UserCreated is the default user created response
	UserCreated = &webserver.Response{
		HTTPStatusCode: http.StatusCreated,
		StatusText:     "201 ",
		Message: "User created successfully.",
	}
)
