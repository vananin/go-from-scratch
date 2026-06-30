package http

import (
	nethttp "net/http"

	"vananin/go-from-scratch/internal/service"
)

func NewRouter(userSvc *service.UserService) nethttp.Handler {
	mux := nethttp.NewServeMux()

	userHandler := NewUserHandler(userSvc)

	mux.HandleFunc("GET /users/{id}", userHandler.GetByID)

	return TimeoutMiddleware(mux)
}
