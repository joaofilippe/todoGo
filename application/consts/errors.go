package consts

import "errors"

var (
	ErrUserAlreadyExists = errors.New("username or email already exists in database")
	ErrInvalidLogin      = errors.New("invalid login")
	ErrUserDoesNotExist  = errors.New("user does not exist")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrSecretKey         = errors.New("secret key not found")
	ErrInvalidNewUser    = errors.New("invalid new user")
)
