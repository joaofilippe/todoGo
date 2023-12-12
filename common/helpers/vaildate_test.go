package helpers_test

import (
	"testing"
	"todoGo/common/helpers"
)

func Test_CheckEmptyFields(t *testing.T) {
	tests := []struct {
		name   string
		data   any
		hasErr    bool
		emptyFields int
	}{
		{
			name:   "empty string",
			data:   "",
			hasErr:    true,
			emptyFields: 0,
		},
		{
			name:   "empty struct",
			data:   struct{}{},
			hasErr:    true,
			emptyFields: 0,
		},
		{
			name: "empty struct with pointer",
			data: struct {
				Pointer *[]string
			}{
			},
			hasErr:    false,
			emptyFields: 1,
		},
		{
			name: "struct with one empty field",
			data: struct {
				name    string
				age     int
			}{
				name:   "João Filippe",
			},
			hasErr:    false,
			emptyFields: 1,
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
