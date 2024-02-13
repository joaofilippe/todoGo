package structs

import (
	"errors"
	"reflect"
	"strings"
)

// GetTagsValues returns the values of the informed tag
func GetTagsValues(s any, tag string)( []string, error) {
	v := reflect.TypeOf(s)

	if !(v.Kind() == reflect.Struct) {
		return []string{}, errors.New("this object isn't a struct")
	}

	tagsValues := []string{}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		value := f.Tag.Get(tag)

		if !strings.EqualFold(value, ""){
			tagsValues = append(tagsValues, )
		}
	}

	return tagsValues, nil
}
