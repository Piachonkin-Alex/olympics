package db

import (
	"database/sql"

	"olympics/pkg/storage"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	pool *sql.DB
}

var (
	psql                 = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	_    storage.Storage = (*Storage)(nil)
)

var _ storage.Storage = (*Storage)(nil)

func NewStorage(cfg *storage.Config) (*Storage, error) {
	connConfig, err := pgx.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, err
	}

	pool := stdlib.OpenDB(*connConfig)
	pool.SetMaxIdleConns(cfg.MaxIdleConns)
	pool.SetMaxOpenConns(cfg.MaxOpenConns)
	pool.SetConnMaxIdleTime(cfg.MaxConnIdleTime)
	pool.SetConnMaxLifetime(cfg.MaxConnLifetime)

	return &Storage{pool: pool}, err
}
