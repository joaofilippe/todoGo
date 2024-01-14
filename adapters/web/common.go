package webserver

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrResponse is the default error response
type ErrResponse struct {
    Err            error `json:"-"`
    HTTPStatusCode int   `json:"-"`

    StatusText string `json:"status"`          
    AppCode    int64  `json:"code,omitempty"`  
    ErrorText  string `json:"error,omitempty"`
}


var (
	// ErrNotFound is the default not found error
    ErrNotFound            = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
    
	// ErrBadRequest is the default bad request error
	ErrBadRequest          = &ErrResponse{HTTPStatusCode: 400, StatusText: "Bad request"}
    
	// ErrInternalServerError is the default internal server error
	ErrInternalServerError = &ErrResponse{HTTPStatusCode: 500, StatusText: "Internal Server Error"}
)

// Render renders the error response
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
    render.Status(r, e.HTTPStatusCode)
    return nil
}