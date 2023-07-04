package usecase

import (
	"context"
	"recipes/domain"
	"recipes/internal/config"
	"recipes/pkg/logger"
)

type storage interface {
	WriteRecipe(ctx context.Context, req domain.Recipe) error
	DeleteRecipe(ctx context.Context, req domain.ID) error
	ListRecipes(ctx context.Context) ([]domain.Recipe, error)
	ReadRecipe(ctx context.Context, req domain.ID) (domain.Recipe, error)
	FindRecipe(ctx context.Context, req domain.Query) ([]domain.Recipe, error)
}

type usecase struct {
	lg   logger.Logger
	stor storage
}

func New(lg logger.Logger, cfg config.Config, stor storage) (*usecase, error) {
	return &usecase{
		lg:   lg,
		stor: stor,
	}, nil
}
