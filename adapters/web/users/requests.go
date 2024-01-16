package users
import "net/http"

// NewUserRequest represents the new user request
type NewUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
}

// Bind binds the request
func (n *NewUserRequest) Bind(r *http.Request) error {
	return nil
}
