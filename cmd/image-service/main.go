package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"vananin/go-from-scratch/internal/app"
	"vananin/go-from-scratch/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	a, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	if err := a.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
