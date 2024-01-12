package webserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Server represents the web server
type Server struct {
	Router *mux.Router
}

// NewServer creates a new server instance
func NewServer() *Server {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	
	return &Server{
		Router: r,
	}
}