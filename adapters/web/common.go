package webserver

import "net/http"

var (
	// ErrNotFound is the default not found error
	ErrNotFound = &ErrResponse{Status: http.StatusNotFound, ErrorMessage: "Resource not found."}

	// ErrBadRequest is the default bad request error
	ErrBadRequest = &ErrResponse{Status: http.StatusBadRequest, ErrorMessage: "Bad request"}

	// ErrInternalServerError is the default internal server error
	ErrInternalServerError = &ErrResponse{Status: http.StatusInternalServerError, ErrorMessage: "Internal Server Error"}
)
