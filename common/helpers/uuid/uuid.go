package helpers

import 	"github.com/google/uuid"

// IsEmpty checks if a uuid is empty
func IsEmpty(id uuid.UUID ) bool {
	return id == uuid.Nil
}