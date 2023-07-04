package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"recipes/domain"
)

func (s *storage) ReadUser(ctx context.Context, login string) (domain.User, error) {
	q01 := `select login, password from users where login = $1`
	row := s.db.QueryRowContext(ctx, q01, login)
	if row.Err() != nil {
		return domain.User{}, fmt.Errorf("tx.s.db.QueryRowContext: %w", row.Err())
	}
	var ret domain.User
	err := row.Scan(&ret.Login, &ret.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, nil
		}
		return domain.User{}, fmt.Errorf("rows.Scan: %w", err)
	}
	return ret, nil
}
