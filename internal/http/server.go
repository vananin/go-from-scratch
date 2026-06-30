package http

import (
	"fmt"
	nethttp "net/http"

	"vananin/go-from-scratch/internal/config"
	"vananin/go-from-scratch/internal/service"
)

func NewServer(cfg config.HTTPConfig, userSvc *service.UserService) *nethttp.Server {
	return &nethttp.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: NewRouter(userSvc),
	}
}
