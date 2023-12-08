package helpers

import (
	"errors"
	"reflect"
)

// ValidateFieldsStruct returns if there is any empty struct field
func ValidateFieldsStruct(s any) ([]string, error) {
	structType := reflect.TypeOf(s)

	if !(structType.Kind() == reflect.Struct) {
		return []string{}, errors.New("this object isn't a struct")
	}

	fieldsNumber := structType.NumField()
	structVal := reflect.ValueOf(s)
	emptyFieldNames := []string{}

	for i := 0; i < fieldsNumber; i++ {
		field := structVal.Field(i)
		fieldName := structType.Field(i).Name

		if field.Kind() == reflect.Pointer {
			if field.IsNil() {
				emptyFieldNames = append(emptyFieldNames, fieldName)
			}
		} else {
			if !(field.IsValid() && !field.IsZero()){
				emptyFieldNames = append(emptyFieldNames, fieldName)
			} 
		}
	}

	return emptyFieldNames, nil
}
