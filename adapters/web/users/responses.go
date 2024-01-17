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
)

var (
	// UserCreated is the default user created response
	UserCreated = &common.Response{
		Status: http.StatusCreated,
		Message: "User created successfully.",
	}
)
