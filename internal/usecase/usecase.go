package usecase

import (
	"context"
	"io"
	"recipes/domain"
	"recipes/pkg/logger"
)

//go:generate mockery --name=storage --structname=storageMock --filename=storage_mock.go --inpackage
type storage interface {
	WriteRecipe(ctx context.Context, req domain.Recipe) error
	DeleteRecipe(ctx context.Context, req domain.ID) error
	ListRecipes(ctx context.Context) ([]domain.RecipeForList, error)
	ReadRecipe(ctx context.Context, req domain.ID) (domain.Recipe, error)
	FindRecipe(ctx context.Context, req domain.Query) ([]domain.Recipe, error)
	VoteRecipe(ctx context.Context, req domain.Vote) error
	ReadUser(ctx context.Context, login string) (domain.User, error)
}

//go:generate mockery --name=filestorage --structname=filestorageMock --filename=filestorage_mock.go --inpackage
type filestorage interface {
	Upload(ctx context.Context, fileName string, fileSize int64, reader io.Reader) error
	Download(ctx context.Context, fileName string) (io.Reader, error)
}

type usecase struct {
	lg   logger.Logger
	stor storage
	fs   filestorage
}

func New(lg logger.Logger, stor storage, fs filestorage) (*usecase, error) {
	return &usecase{
		lg:   lg,
		stor: stor,
		fs:   fs,
	}, nil
}
