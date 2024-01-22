package users

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"

	"github.com/joaofilippe/todoGo/adapters/web/common"
	dto "github.com/joaofilippe/todoGo/adapters/web/dto/users"
	"github.com/joaofilippe/todoGo/application"
	uuid_helper "github.com/joaofilippe/todoGo/pkg/uuid"
)

// Web represents the user web adapter
type Web struct {
	Application *application.Application
}

// Create creates a new user
func (u *Web) Create(w http.ResponseWriter, r *http.Request) {
	n := new(NewUserRequest)
	if err := render.Bind(r, n); err != nil {
		render.Render(w, r, common.ErrBadRequest)
		return
	}

	dto := new(NewUserDTO)

	dto.FromRequestToDTO(*n)

	newUser, err := dto.FromDTOToModel()
	if err != nil {
		render.Render(w, r, ErrCannotCreateUser)
		return
	}

	id, err := u.Application.UserService.Create(newUser)
	if err != nil {
		e := ErrCannotCreateUser
		e.Data = err.Error()
		render.Render(w, r, ErrCannotCreateUser)
		return
	}

	if uuid_helper.IsEmpty(id) {
		render.Render(w, r, ErrCannotCreateUser)
		return
	}

	response := UserCreated
	response.Data = struct {
		UserID uuid.UUID `json:"user_id"`
	}{
		UserID: id,
	}

	render.Render(w, r, response)
}

// GetUserByUsername gets a user by username
func (u *Web) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	user, err := u.Application.UserService.GetUserByUsername(username)
	if err != nil {
		render.Render(w, r, ErrCannotGetUser)
		return
	}

	response := GetUserResponse
	response.Data = user

	render.Render(w, r, response)
}

// GetUserByFilter gets a user by filter
func (u *Web) GetUserByFilter(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")

	var users []dto.User
	if username != "" {
		user, err := u.Application.UserService.GetUserByUsername(username)
		if err != nil {
			render.Render(w, r, ErrCannotGetUser)
			return
		}

		if user.ID != uuid.Nil {
			users = append(users, *dto.FromDomain(user))
		}
	}

	if email != "" {
		user, err := u.Application.UserService.GetUserByEmail(email)
		if err != nil {
			render.Render(w, r, ErrCannotGetUser)
			return
		}

		if user.ID != uuid.Nil {
			if len(users) > 0 {
				for _, u := range users {
					if u.ID != user.ID.String() {
						users = append(users, *dto.FromDomain(user))
					}
				}
			} else {
				users = append(users, *dto.FromDomain(user))
			}
		}
	}

	if len(users) == 0 {
		render.Render(w, r, ErrUserNotFound)
		return
	}

	response := GetUserResponse
	response.Data = struct {
		User interface{} `json:"users"`
	}{
		User: users,
	}

	render.Render(w, r, response)
}

// GetUserByID gets a user by id
func (u *Web) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		render.Render(w, r, ErrInvalidID)
		return
	}

	user, err := u.Application.UserService.GetUserByID(id)
	if err != nil {
		render.Render(w, r, ErrCannotGetUser)
		return
	}

	response := GetUserResponse
	response.Data = user

	render.Render(w, r, response)
}

// GetUserByEmail gets a user by email
func (u *Web) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	user, err := u.Application.UserService.GetUserByEmail(email)
	if err != nil {
		render.Render(w, r, ErrCannotGetUser)
		return
	}

	response := GetUserResponse
	response.Data = user

	render.Render(w, r, response)
}
