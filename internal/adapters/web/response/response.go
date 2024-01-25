package response

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Status int `json:"-"`

	StatusText string      `json:"status"`
	AppCode    int64       `json:"code,omitempty"`
	Message    string      `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func (r *Response) Render(w http.ResponseWriter, req *http.Request) error {
	render.Status(req, r.Status)
	return nil
}
