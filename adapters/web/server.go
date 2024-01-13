package webserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaofilippe/todoGo/application"
)

// Server represents the web server
type Server struct {
	Application *application.Application
	Router *mux.Router
}

// NewServer creates a new server instance
func NewServer(application *application.Application) *Server {
	r := mux.NewRouter()

	fmt.Print("Creating new server instance...")

	r.HandleFunc("/", homeHandler)
	
	return &Server{
		Application: application,
		Router: r,
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it's alive!")
}