package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"recipes/domain"
	"recipes/pkg/logger"

	"github.com/redis/go-redis/v9"
)

type Handler struct {
	lg logger.Logger
	uc useCase
	rc *redis.Client
}

type useCase interface {
	CreateRecipe(ctx context.Context, req domain.Recipe) (domain.Recipe, error)
	UpdateRecipe(ctx context.Context, req domain.Recipe) (domain.Recipe, error)
	DeleteRecipe(ctx context.Context, req domain.ID) (domain.ID, error)
	ListRecipes(ctx context.Context) ([]domain.RecipeForList, error)
	ReadRecipe(ctx context.Context, req domain.ID) (domain.Recipe, error)
	FindRecipe(ctx context.Context, req domain.Query) ([]domain.Recipe, error)
	VoteRecipe(ctx context.Context, req domain.Vote) error
	UploadRecipe(ctx context.Context, req domain.FileInfoUpload) error
	DownloadRecipe(ctx context.Context, req domain.FileInfoDownload) (io.Reader, error)
	SignIn(ctx context.Context, req domain.User) (domain.User, bool, error)
}

func New(lg logger.Logger, useCase useCase, rcli *redis.Client) *Handler {
	return &Handler{
		lg: lg,
		uc: useCase,
		rc: rcli,
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
