package user

import "errors"

var (
	// ErrNoUsername is an error that occurs when the username is empty
	ErrNoUsername = errors.New("username is empty")

	// ErrNoEmail is an error that occurs when the email is empty
	ErrNoEmail = errors.New("email is empty")

	// ErrNoID is an error that occurs when the id is empty
	ErrNoID = errors.New("id is empty")
)
