package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"recipes/domain"
	"time"

	"github.com/lib/pq"
)

func (s *storage) WriteRecipe(ctx context.Context, req domain.Recipe) error {
	q01 := `insert into recipes (id, cr_dt, del_dt, user_id, title, description,
		ingredients, steps, total_time)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		on conflict (id) do update
	set upd_dt = $2, del_dt = $3, user_id = $4, title = $5, description = $6,
		ingredients = $7, steps = $8, total_time = $9`

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("tx.Begin: %w", err)
	}
	defer tx.Rollback()

	steps, err := json.Marshal(req.Steps)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}
	_, err = tx.ExecContext(ctx, q01, req.Id, req.CrDt, req.DelDt, req.UserId, req.Title,
		req.Description, pq.Array(req.Ingredients), steps, req.TotalTime)
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

func (s *storage) ListRecipes(ctx context.Context) ([]domain.RecipeForList, error) {
	q01 := `select id, title from recipes where del_dt is null`

	rows, err := s.db.QueryContext(ctx, q01)
	if err != nil {
		return nil, fmt.Errorf("s.db.QueryContext: %w", err)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("s.db.QueryContext: %w", rows.Err())
	}
	var ret []domain.RecipeForList
	for rows.Next() {
		var item domain.RecipeForList
		err = rows.Scan(&item.Id, &item.Title)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		ret = append(ret, item)
	}
	return ret, nil
}

func (s *storage) ReadRecipe(ctx context.Context, req domain.ID) (domain.Recipe, error) {
	q01 := `select id, title, description, ingredients, steps, total_time, rating
				from recipes
			where id = $1 and del_dt is null`

	row := s.db.QueryRowContext(ctx, q01, req.Id)
	if row.Err() != nil {
		return domain.Recipe{}, fmt.Errorf("tx.s.db.QueryRowContext: %w", row.Err())
	}
	var ret domain.Recipe
	var steps []byte
	err := row.Scan(&ret.Id, &ret.Title, &ret.Description, pq.Array(&ret.Ingredients),
		&steps, &ret.TotalTime, &ret.Rating)
	if err != nil {
		return domain.Recipe{}, fmt.Errorf("rows.Scan: %w", err)
	}
	err = json.Unmarshal(steps, &ret.Steps)
	if err != nil {
		return domain.Recipe{}, fmt.Errorf("json.Unmarshal: %w", err)
	}
	return ret, nil
}

func (s *storage) FindRecipe(ctx context.Context, req domain.Query) ([]domain.Recipe, error) {
	q01 := `select id, title, description, ingredients, steps, total_time, rating from recipes
				where $1 <@ ingredients and del_dt is null`
	q02 := ` and total_time <= $2`
	q05 := ` and rating >= $3`

	var rows *sql.Rows
	var err error
	q04 := q01
	if req.MaxTime > 0 {
		q04 += q02
	}
	if req.MinRating > 0 {
		q04 += q05
	}

	if req.SortByTime != "" || req.SortByRating != "" {
		q04 += " order by "
		if req.SortByTime != "" {
			q04 += "total_time " + req.SortByTime
			if req.SortByRating != "" {
				q04 += ", "
			}
		}
		if req.SortByTime != "" {
			q04 += "rating " + req.SortByRating
		}
	}
	if req.MaxTime > 0 && req.MinRating > 0 {
		rows, err = s.db.QueryContext(ctx, q04, pq.Array(req.Ingredients), req.MaxTime, req.MinRating)
	} else if req.MaxTime > 0 && req.MinRating == 0 {
		rows, err = s.db.QueryContext(ctx, q04, pq.Array(req.Ingredients), req.MaxTime)
	} else if req.MaxTime == 0 && req.MinRating > 0 {
		rows, err = s.db.QueryContext(ctx, q04, pq.Array(req.Ingredients), req.MinRating)
	} else {
		rows, err = s.db.QueryContext(ctx, q04, pq.Array(req.Ingredients))
	}
	if err != nil {
		return nil, fmt.Errorf("s.db.QueryContext: %w", err)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("s.db.QueryContext: %w", rows.Err())
	}
	var ret []domain.Recipe
	var steps []byte
	for rows.Next() {
		var item domain.Recipe
		err = rows.Scan(&item.Id, &item.Title, &item.Description, pq.Array(&item.Ingredients),
			&steps, &item.TotalTime, &item.Rating)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		err = json.Unmarshal(steps, &item.Steps)
		if err != nil {
			return nil, fmt.Errorf("json.Unmarshal: %w", err)
		}
		ret = append(ret, item)
	}
	return ret, nil
}

func (s *storage) VoteRecipe(ctx context.Context, req domain.Vote) error {
	q01 := `select mark from votes where recipe_id = $1`
	rows, err := s.db.QueryContext(ctx, q01, req.RecipeId)
	if err != nil {
		return fmt.Errorf("s.db.QueryContext: %w", err)
	}
	if rows.Err() != nil {
		return fmt.Errorf("rows.Err(): %w", err)
	}
	defer rows.Close()
	cnt := 1
	total := req.Mark
	for rows.Next() {
		var mark int
		err = rows.Scan(&mark)
		if err != nil {
			return fmt.Errorf("rows.Scan: %w", err)
		}
		total += mark
		cnt++
	}
	rating := float64(total) / float64(cnt)
	rating = math.Floor(rating*100) / 100

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("tx.Begin: %w", err)
	}
	defer tx.Rollback()

	q02 := `insert into votes (cr_dt, recipe_id, user_id, mark) values ($1, $2, $3, $4)`
	_, err = tx.ExecContext(ctx, q02, req.CrDt, req.RecipeId, req.UserId, req.Mark)
	if err != nil {
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code == "23505" {
			return domain.ErrDuplicateRecord
		}
		return fmt.Errorf("tx.ExecContext: %w", err)
	}

	q03 := `update recipes set rating = $1 where id = $2`
	_, err = tx.ExecContext(ctx, q03, rating, req.RecipeId)
	if err != nil {
		return fmt.Errorf("tx.ExecContext: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("tx.Commit: %w", err)
	}
	return nil
}
