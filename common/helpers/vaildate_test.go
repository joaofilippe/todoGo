package helpers_test

import (
	"testing"
	"github.com/joaofilippe/todoGo/common/helpers"
)

func Test_CheckEmptyFields(t *testing.T) {
	tests := []struct {
		name        string
		data        any
		hasErr      bool
		emptyFields int
	}{
		{
			name:        "test 1 -empty struct",
			data:        struct{}{},
			hasErr:      true,
			emptyFields: 0,
		},
		{
			name:        "test 2 - empty string",
			data:        "",
			hasErr:      true,
			emptyFields: 0,
		},
		{
			name: "test 3 - empty struct with pointer",
			data: struct {
				Pointer *[]string
			}{},
			hasErr:      true,
			emptyFields: 1,
		},
		{
			name: "test 4 - struct with one empty field",
			data: struct {
				name string
				age  int
			}{
				name: "João Filippe",
			},
			hasErr:      true,
			emptyFields: 1,
		},
		{
			name: "test 5 - struct with two empty field",
			data: struct {
				name  string
				age   int
				email string
			}{
				name: "João Filippe",
			},
			hasErr:      true,
			emptyFields: 2,
		},
		{
			name: "struct with two empty field, but one is a pointer",
			data: struct {
				name        string
				age         int
				email       string
				stringSlice *[]string
			}{
				name: "João Filippe",
				age: 30,
			},
			hasErr:      true,
			emptyFields: 2,
		},
		{
			name: "struct with no empty field",
			data: struct {
				name  string
				age   int
				email string
			}{
				name: "João Filippe",
				age : 30,
				email: "joaofilippe@outlook.com",
			},
			hasErr:      false,
			emptyFields: 0,
		},
	}

	for _, tt := range tests {
		res, err := helpers.CheckEmptyFields(tt.data)
		if err != nil && !tt.hasErr {
			t.Error(err)
		}

		if len(res) != tt.emptyFields {
			t.Errorf("expected %d empty fields, got %d", tt.emptyFields, len(res))
		}
	}
}
