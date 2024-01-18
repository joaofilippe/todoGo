package helpers

import (
	"strings"

	"github.com/google/uuid"
)

// IsEmpty checks if a uuid is empty
func IsEmpty(id uuid.UUID) bool {
	return id == uuid.Nil
}

// ConvertStringToArray converts a string to an array of uuids
func ConvertStringToArray(str string) ([]uuid.UUID, error) {
	var uuids []uuid.UUID
	
	if str == "" {
		return uuids, nil
	}

	strArray := strings.Split(str, ",")

	for i := range strArray {
		uuid, err := uuid.Parse(strArray[i])
		if err != nil {
			return nil, err
		}

		uuids = append(uuids, uuid)
	}

	return uuids, nil
}
