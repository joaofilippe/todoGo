package user

import (
	"net/http"

	webserver "github.com/joaofilippe/todoGo/adapters/web"
)

var (
	// ErrCannotCreateUser is the default cannot create user error
	ErrCannotCreateUser = &webserver.ErrResponse{Status: http.StatusInternalServerError, ErrorMessage: "Cannot create user."}

	// ErrUserNotFound is the default user not found error
	ErrUserNotFound = &webserver.Response{Status: http.StatusNoContent, Message: "User not found."}
)

var (
	// UserCreated is the default user created response
	UserCreated = &webserver.Response{
		Status: http.StatusCreated,
		Message: "User created successfully.",
	}
)
