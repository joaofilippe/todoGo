package email

import "testing"

func Test_Email(t *testing.T){
	email := "joaofilippeoutlook"

	if !Validate(email) {
		t.Error("Email is not valid")
	}
}