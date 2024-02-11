package structs

import (
	"errors"
	"reflect"
)

func GetTagsValues(s any, tag string)( []string, error) {
	v := reflect.TypeOf(s)

	if !(v.Kind() == reflect.Struct) {
		return []string{}, errors.New("this object isn't a struct")
	}

	tagsValues := []string{}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		tagsValues = append(tagsValues, f.Tag.Get(tag))
	}

	return tagsValues, nil
}
