package webserver

import (
	"net/http"

	"github.com/go-chi/render"
)

// Response represents the default response
type Response struct {
	HTTPStatusCode int `json:"-"`

	StatusText string      `json:"status"`
	AppCode    int64       `json:"code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// Render renders the response
func (r *Response) Render(w http.ResponseWriter, req *http.Request) error {
	render.Status(req, r.HTTPStatusCode)
	return nil
}

// ErrResponse is the default error response
type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText   string      `json:"status"`
	AppCode      int64       `json:"code,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
