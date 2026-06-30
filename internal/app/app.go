package app

import (
	"database/sql"
	"net/http"
	"vananin/go-from-scratch/internal/config"
	"vananin/go-from-scratch/internal/db"
	httpSrv "vananin/go-from-scratch/internal/http"
	"vananin/go-from-scratch/internal/repository/user"
	"vananin/go-from-scratch/internal/service"
)

type App struct {
	db *sql.DB

	httpServer *http.Server
	// grpcServer *grpc.Server
}

func New(cfg config.Config) (*App, error) {
	db, err := db.NewPostgres(cfg.DB)
	if err != nil {
		return nil, err
	}

	userRepo := user.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)

	httpServer := httpSrv.NewServer(cfg.HTTP, userSvc)
	// grpcServer := grpc.NewServer(userSvc)

	return &App{
		db:         db,
		httpServer: httpServer,
		// grpcServer: grpcServer,
	}, nil
}
