package webserver

import "github.com/gorilla/mux"

// Server represents the web server
type Server struct {
	Router *mux.Router
}

// NewServer creates a new server instance
func NewServer() *Server {
	r := mux.NewRouter()

	return &Server{
		Router: r,
	}
}