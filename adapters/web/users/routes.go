package users

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func (u *Web) BuildRoutes(router chi.Router) {
	router.Route("/user", func(r chi.Router) {
		r.Get("/", testUserRoute)
		r.Post("/", u.Create)
	})
}

func testUserRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "user route!")
}