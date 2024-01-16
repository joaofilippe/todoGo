package email

import (
	"fmt"
	"net/mail"
)

func ValidateEmail(email string) bool {
	e, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	fmt.Println(e)

	return true
}