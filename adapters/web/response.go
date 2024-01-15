package webserver

import (
	"net/http"

	"github.com/go-chi/render"
)

// Response represents the default response
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Render renders the response
func (r *Response) Render(w http.ResponseWriter, req *http.Request) error {
	render.Status(req, r.Status)
	return nil
}

// ErrResponse is the default error response
type ErrResponse struct {
	Status       int         `json:"status"`
	ErrorMessage string      `json:"error_message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

// Render renders the error response
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Status)
	return nil
}
