package helpers_test

import (
	"fmt"
	"testing"
	"todoGo/common/helpers"
)

func Test_CheckEmptyFields(t *testing.T) {
	tests := []struct {
		name string
		data any
		want bool
	}{
		{
			name: "empty string",
			data: "",
			want: true,
		},
		{
			name: "empty struct",
			data: struct{}{},
			want: true,
		},
		{
			name: "empty struct with pointer",
			data: struct {
				Pointer *string
			}{},
		},
		{
			name: "struct with one empty field",
			data: struct {
				name string
				Pointer *string
			}{},
		},
	}

	for _, tt := range tests {
		res, err := helpers.CheckEmptyFields(tt.data)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(res)
	}
}
