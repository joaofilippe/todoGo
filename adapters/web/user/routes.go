package user

import (
	"net/http"

	"github.com/go-chi/render"

	webserver "github.com/joaofilippe/todoGo/adapters/web"
	"github.com/joaofilippe/todoGo/application"
	models "github.com/joaofilippe/todoGo/application/models/user"
)

// Web represents the user web adapter
type Web struct {
	Application *application.Application
	WebServer   *webserver.Server
}

type UserRequest struct {

}

func (u *UserRequest) Bind(r *http.Request) error {
	return nil
}

// Create creates a new user
func (u *Web) Create(w http.ResponseWriter, r *http.Request) {
	n := new(UserRequest)
	if err := render.Bind(r, n); err != nil {
		render.Render(w, r, webserver.ErrBadRequest)
		return
	}

	newUser := new(models.NewUser)

	u.Application.UserService.CreateUser(newUser)
}
