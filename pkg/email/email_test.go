package email

import "testing"

func Test_Email(t *testing.T){
	email := "joaofilippe@outlook"

	if !ValidateEmail(email) {
		t.Error("Email is not valid")
	}
}