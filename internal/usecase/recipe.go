package usecase

import (
	"context"
	"fmt"
	"io"
	"recipes/domain"
	"recipes/pkg/tools"
)

func (u *usecase) CreateRecipe(ctx context.Context, req domain.Recipe) (domain.Recipe, error) {
	req.ID.Id = tools.GetGuid()
	err := u.stor.WriteRecipe(ctx, req)
	if err != nil {
		return domain.Recipe{}, fmt.Errorf("u.stor.WriteRecipe: %w", err)
	}
	return req, nil
}

func (u *usecase) UpdateRecipe(ctx context.Context, req domain.Recipe) (domain.Recipe, error) {
	err := u.stor.WriteRecipe(ctx, req)
	if err != nil {
		return domain.Recipe{}, fmt.Errorf("u.stor.WriteRecipe: %w", err)
	}
	return req, nil
}

func (u *usecase) DeleteRecipe(ctx context.Context, req domain.ID) (domain.ID, error) {
	err := u.stor.DeleteRecipe(ctx, req)
	if err != nil {
		return domain.ID{}, fmt.Errorf("u.stor.DeleteRecipe: %w", err)
	}
	return req, nil
}

func (u *usecase) ListRecipes(ctx context.Context) ([]domain.RecipeForList, error) {
	return u.stor.ListRecipes(ctx)
}

func (u *usecase) ReadRecipe(ctx context.Context, req domain.ID) (domain.Recipe, error) {
	return u.stor.ReadRecipe(ctx, req)
}

func (u *usecase) FindRecipe(ctx context.Context, req domain.Query) ([]domain.Recipe, error) {
	return u.stor.FindRecipe(ctx, req)
}

func (u *usecase) VoteRecipe(ctx context.Context, req domain.Vote) error {
	return u.stor.VoteRecipe(ctx, req)
}

func (u *usecase) UploadRecipe(ctx context.Context, req domain.FileInfoUpload) error {
	err := u.fs.Upload(ctx, req.Id+"_"+req.Step, req.Size, req.Reader)
	if err != nil {
		return fmt.Errorf("u.fs.Upload: %w", err)
	}
	return nil
}

func (u *usecase) DownloadRecipe(ctx context.Context, req domain.FileInfoDownload) (io.Reader, error) {
	reader, err := u.fs.Download(ctx, req.Id+"_"+req.Step)
	if err != nil {
		return nil, fmt.Errorf("u.fs.Download: %w", err)
	}
	return reader, nil
}
