package usecase

import (
	"context"
	"fmt"
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

func (u *usecase) ListRecipes(ctx context.Context) ([]domain.Recipe, error) {
	return u.stor.ListRecipes(ctx)
}

func (u *usecase) ReadRecipe(ctx context.Context, req domain.ID) (domain.Recipe, error) {
	return u.stor.ReadRecipe(ctx, req)
}
