package http

import (
	"encoding/json"
	nethttp "net/http"
)

type Responder struct{}

func NewResponder() *Responder {
	return &Responder{}
}

func (r *Responder) JSON(w nethttp.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func (r *Responder) Error(w nethttp.ResponseWriter, status int, message string) error {
	return r.JSON(w, status, map[string]string{
		"error": message,
	})
}
