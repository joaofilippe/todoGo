package structs

import (
	"errors"
	"reflect"
	"strings"
)

// GetTagsValues returns the values of the informed tag
func GetTagsValues(s any, tag string) ([]string, error) {
	t := reflect.TypeOf(s)

	if !(t.Kind() == reflect.Struct) {
		return []string{}, errors.New("this object isn't a struct")
	}

	tagsValues := []string{}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		value := f.Tag.Get(tag)

		if !strings.EqualFold(value, "") {
			tagsValues = append(tagsValues)
		}
	}

	return tagsValues, nil
}

// SetStringFieldValueByName set value toa
func SetStringFieldValueByName(s any, fieldName string, value string) error {
	e := reflect.ValueOf(s).Elem()

	if !(e.Kind() == reflect.Struct) {
		return errors.New("this object isn't a struct")
	}

	field := e.FieldByName(fieldName)
	if field == (reflect.Value{}) {
		return errors.New("field doesn't exist")
	}

	if not := field.CanSet(); !not {
		return errors.New("field cannot be ")
	}

	field.Set(reflect.ValueOf(value))

	return nil
}
