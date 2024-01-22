package users

import (
	"net/http"

	"github.com/joaofilippe/todoGo/adapters/web/common"
)

var (
	// ErrCannotCreateUser is the default cannot create user error
	ErrCannotCreateUser = &common.ErrResponse{Status: http.StatusInternalServerError, ErrorMessage: "Cannot create user."}

	// ErrUserNotFound is the default user not found error
	ErrUserNotFound = &common.Response{Status: http.StatusNoContent, Message: "User not found."}

	// ErrCannotConvertUser is the default cannot convert user error
	ErrCannotConvertUser = &common.Response{Status: http.StatusInternalServerError, Message: "Cannot convert user dto."}

	// ErrInvalidID is the default invalid id error
	ErrInvalidID = &common.Response{Status: http.StatusBadRequest, Message: "Invalid id."}

	// ErrCannotGetUser is the default cannot get user error
	ErrCannotGetUser = &common.Response{Status: http.StatusInternalServerError, Message: "Cannot get user."}
)

var (
	// UserCreated is the default user created response
	UserCreated = &common.Response{
		Status:  http.StatusCreated,
		Message: "User created successfully.",
	}

	// GetUserResponse is the default get user response
	GetUserResponse = &common.Response{
		Status:  http.StatusOK,
		Message: "Users retrieved successfully.",
	}

	// UserNotFoundResponse is the default get user not exists response
	UserNotFoundResponse = &common.Response{
		Status:  http.StatusNoContent,
	}
)
