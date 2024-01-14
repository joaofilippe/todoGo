package webserver

import (
	"fmt"
	"net/http"
	"os"

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

// Run starts the server
func (s *Server) Run() error {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port :"+ port)

	return http.ListenAndServe(":"+port, s.Router)
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)

	ok := `
		<h1>Ok</h1>
		<p>It's working!</p>
	`

	
	w.Write([]byte(ok))
}
