package structs

import(
	"github.com/stretchr/testify/assert"
	"testing"
	) 
	
func Test_GetTags(t *testing.T){
	type St struct {
		Field  string `tag:"valor"`
		Field2 string `tag:"valor2"`
	}

	s := St{
		Field: "field",
		Field2: "field_2",
	}

	values, err := GetTagsValues(s, "tag")
	assert.NotEmpty(t, len(values), "no values returned")
	assert.Nil(t, err, "error must be nil")
}