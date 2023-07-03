package main

import (
	"context"
	"log"
	"os/signal"
	"recipes/internal/app"
	"recipes/internal/config"
	"recipes/pkg/logger"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	lg, err := logger.New(true)
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.New()
	if err != nil {
		lg.Fatal(err.Error())
	}

	a, err := app.New(lg, cfg)
	if err != nil {
		lg.Fatal(err.Error())
	}

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return a.Run(ctx)
	})
	if err = eg.Wait(); err != nil {
		lg.Fatal(err.Error())
	}
}
