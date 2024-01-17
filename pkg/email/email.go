package email

import (
	"fmt"
	"net/mail"
)

// ValidateEmail validates an email
func Validate(email string) bool {
	e, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	fmt.Println(e)

	return true
}