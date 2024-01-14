package webserver

import "net/http"

var (
	// ErrNotFound is the default not found error
	ErrNotFound = &ErrResponse{HTTPStatusCode: http.StatusNotFound, StatusText: "Resource not found."}

	// ErrBadRequest is the default bad request error
	ErrBadRequest = &ErrResponse{HTTPStatusCode: http.StatusBadRequest, StatusText: "Bad request"}

	// ErrInternalServerError is the default internal server error
	ErrInternalServerError = &ErrResponse{HTTPStatusCode: http.StatusInternalServerError, StatusText: "Internal Server Error"}
)
