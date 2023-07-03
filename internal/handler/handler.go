package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"recipes/domain"
	"recipes/pkg/logger"
)

type Handler struct {
	lg logger.Logger
	uc useCase
}

type useCase interface {
	CreateRecipe(ctx context.Context, req domain.Recipe) (domain.Recipe, error)
	UpdateRecipe(ctx context.Context, req domain.Recipe) (domain.Recipe, error)
	DeleteRecipe(ctx context.Context, req domain.ID) (domain.ID, error)
	ListRecipes(ctx context.Context) ([]domain.Recipe, error)
	ReadRecipe(ctx context.Context, req domain.ID) (domain.Recipe, error)
}

func New(lg logger.Logger, useCase useCase) *Handler {
	return &Handler{
		lg: lg,
		uc: useCase,
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
