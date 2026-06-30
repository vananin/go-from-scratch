package app

import (
	"context"
	"errors"
)

func (a *App) Shutdown(ctx context.Context) error {
	var errs []error

	if err := a.httpServer.Shutdown(ctx); err != nil {
		errs = append(errs, err)
	}

	// a.grpcServer.GracefulStop()

	if err := a.db.Close(); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}
