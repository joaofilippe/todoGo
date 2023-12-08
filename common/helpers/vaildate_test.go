package helpers_test

import (
	"fmt"
	"testing"
	"todoGo/common/helpers"
)

func Test_ValidateFieldsStruct(t *testing.T) {
	s := struct {
		id       int
		name     string
		username string
		password string
		field    string
		pointer *string
	}{
		1,
		"João Filippe",
		"joaofilippe",
		"",
		"",
		nil,
	}

	r, err := helpers.ValidateFieldsStruct(s)
	if err != nil {
		t.Error(err)
	}

	if len(r) > 0 {
		err = fmt.Errorf("the fields is empty: ")
		for i := range r {
			err = fmt.Errorf("%w %s", err, r[i])
		}

		t.Error(err)
	}
}
