package storage

import (
	"context"
	"fmt"
	"recipes/domain"
	"time"

	"github.com/lib/pq"
)

func (s *storage) WriteRecipe(ctx context.Context, req domain.Recipe) error {
	q01 := `insert into recipes (id, cr_dt, del_dt, user_id, title, description,
		ingredients, steps)
	values ($1, $2, $3, $4, $5, $6, $7, $8)
		on conflict (id) do update
	set upd_dt = $2, del_dt = $3, user_id = $4, title = $5, description = $6,
		ingredients = $7, steps = $8`

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("tx.Begin: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, q01, req.Id, req.CrDt, req.DelDt, req.UserId, req.Title,
		req.Description, pq.Array(req.Ingredients), pq.Array(req.Steps))
	if err != nil {
		return fmt.Errorf("tx.ExecContext: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("tx.Commit: %w", err)
	}
	return nil
}

func (s *storage) DeleteRecipe(ctx context.Context, req domain.ID) error {
	q01 := `update recipes set del_dt = $1 where id = $2`

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("tx.Begin: %w", err)
	}
	defer tx.Rollback()

	delDt := time.Now()
	_, err = tx.ExecContext(ctx, q01, delDt, req.Id)
	if err != nil {
		return fmt.Errorf("tx.ExecContext: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("tx.Commit: %w", err)
	}

	return nil
}

func (s *storage) ListRecipes(ctx context.Context) ([]domain.Recipe, error) {
	sql1 := `select id, title, description, ingredients, steps from recipes
				where del_dt is null`

	rows, err := s.db.QueryContext(ctx, sql1)
	if err != nil {
		return nil, fmt.Errorf("tx.s.db.QueryContext: %w", err)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("tx.s.db.QueryContext: %w", rows.Err())
	}
	var ret []domain.Recipe
	for rows.Next() {
		var item domain.Recipe
		err = rows.Scan(&item.Id, &item.Title, &item.Description, pq.Array(&item.Ingredients),
			pq.Array(&item.Steps))
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		ret = append(ret, item)
	}
	return ret, nil
}

func (s *storage) ReadRecipe(ctx context.Context, req domain.ID) (domain.Recipe, error) {
	sql1 := `select id, title, description, ingredients, steps from recipes
				where id = $1 and del_dt is null`

	row := s.db.QueryRowContext(ctx, sql1, req.Id)
	if row.Err() != nil {
		return domain.Recipe{}, fmt.Errorf("tx.s.db.QueryRowContext: %w", row.Err())
	}
	var ret domain.Recipe
	err := row.Scan(&ret.Id, &ret.Title, &ret.Description, pq.Array(&ret.Ingredients),
		pq.Array(&ret.Steps))
	if err != nil {
		return domain.Recipe{}, fmt.Errorf("rows.Scan: %w", err)
	}
	return ret, nil
}
