package users
import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"

	uuid_helper "github.com/joaofilippe/todoGo/common/helpers/uuid"
	"github.com/joaofilippe/todoGo/adapters/web/common"
	"github.com/joaofilippe/todoGo/application"
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

	dto.FromRequestToDTO(n)

	newUser := dto.FromDTOToModel()

	id, err := u.Application.UserService.CreateUser(newUser)
	if err != nil {
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
