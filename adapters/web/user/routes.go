package user

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"

	webserver "github.com/joaofilippe/todoGo/adapters/web"
	"github.com/joaofilippe/todoGo/application"
	models "github.com/joaofilippe/todoGo/application/models/user"
)

// Web represents the user web adapter
type Web struct {
	Application *application.Application
	WebServer   *webserver.Server
}

// Create creates a new user
func (u *Web) Create(w http.ResponseWriter, r *http.Request) {
	n := new(NewUserRequest)
	if err := render.Bind(r, n); err != nil {
		render.Render(w, r, webserver.ErrBadRequest)
		return
	}

	newUser := new(models.NewUser)

	id, err := u.Application.UserService.CreateUser(newUser)
	if err != nil {
		render.Render(w, r, ErrCannotCreateUser)
		return
	}

	if id.

	response := UserCreated
	response.Data = struct {
		UserID uuid.UUID `json:"user_id"`
	}{
		UserID: id,
	}

	render.Render(w, r, response)
}
