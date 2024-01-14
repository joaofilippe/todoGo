package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joaofilippe/todoGo/application"
)

// Server represents the web server
type Server struct {
	Application *application.Application
	Router      *chi.Mux
}

// NewServer creates a new server instance
func NewServer(application *application.Application) *Server {
	r := chi.NewRouter()

	fmt.Print("Creating new server instance...")

	r.HandleFunc("/", homeHandler)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/", apiV1Handler)
		r.Get("/todos", todosV1Handler)
	})

	r.HandleFunc("/todos", todosHandler)
	
	return &Server{
		Application: application,
		Router:      r,
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it's alive!")
}

func apiV1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "api v1")
}

func todosV1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "todos v1")
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "todos")
}
