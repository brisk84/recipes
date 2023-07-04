package app

import (
	"context"
	"fmt"
	"recipes/internal/config"
	"recipes/internal/handler"
	"recipes/internal/server"
	"recipes/internal/storage"
	"recipes/internal/usecase"
	"recipes/pkg/logger"

	"github.com/redis/go-redis/v9"
)

type App struct {
	srv *server.Server
	lg  logger.Logger
}

func New(lg logger.Logger, cfg config.Config) (*App, error) {
	stor, err := storage.New(lg, cfg)
	if err != nil {
		return nil, fmt.Errorf("storage: %w", err)
	}
	uc, err := usecase.New(lg, cfg, stor)
	if err != nil {
		return nil, fmt.Errorf("usecase: %w", err)
	}

	rcli := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	h := handler.New(lg, uc, rcli)
	srv := server.New(lg, cfg.AppAddr, h)

	return &App{
		srv: srv,
		lg:  lg,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.srv.Start(ctx)
}
