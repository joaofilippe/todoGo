package users

import (
	"github.com/go-chi/chi"
)

func (u *Web) BuildRoutes(router chi.Router) {
	router.Route("/user", func(r chi.Router) {
		r.Post("/", u.Create)
 		r.Get("/", u.GetUserByFilter)
		r.Get("/{id}", u.GetUserByID)
	})
}