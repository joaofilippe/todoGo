package webserver

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joaofilippe/todoGo/adapters/web/users"
	"github.com/joaofilippe/todoGo/application"
)

// Server represents the web server
type Server struct {
	Application *application.Application
	User        *users.Web
	Router      *chi.Mux
}

// NewServer creates a new server instance
func NewServer(application *application.Application) *Server {
	r := chi.NewRouter()

	server := &Server{
		Application: application,
		Router:      r,
	}

	server.buildAPIV1Routes()

	return server
}

// Run starts the server
func (s *Server) Run() error {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port :" + port)

	return http.ListenAndServe(":"+port, s.Router)
}

func (s *Server) buildAPIV1Routes() {
	apiV1 := s.Router.Route("/v1", func(r chi.Router) {})

	s.newUserWeb(s.Application, apiV1)

	s.User.BuildRoutes(s.Router)
}

func (s *Server) newUserWeb(application *application.Application, router chi.Router) {
	userWeb := &users.Web{
		Application: application,
	}

	userWeb.BuildRoutes(router)

	s.User = userWeb
}
