package storage

import (
	"database/sql"
	"embed"
	"recipes/internal/config"
	"recipes/pkg/logger"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type storage struct {
	lg    logger.Logger
	db    *sql.DB
	pgURI string
}

func New(logger logger.Logger, cfg config.Config) (*storage, error) {
	stor := storage{
		lg:    logger,
		pgURI: cfg.PgURI,
	}
	if err := stor.Connect(); err != nil {
		return nil, err
	}
	return &stor, nil
}

func (s *storage) Connect() error {
	var err error
	if s.db, err = sql.Open("postgres", s.pgURI); err != nil {
		return err
	}
	if err = s.db.Ping(); err != nil {
		return err
	}
	goose.SetBaseFS(embedMigrations)
	if err := goose.Up(s.db, "migrations"); err != nil {
		return err
	}
	return nil
}

func (s *storage) Close() error {
	err1 := s.db.Close()
	if err1 != nil {
		return err1
	}
	return nil
}
