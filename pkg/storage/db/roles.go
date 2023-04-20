package db

import (
	"context"
	"database/sql"
	"errors"

	"olympics/pkg/core/entities"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) GetInfoByClient(ctx context.Context, name string) (entities.Role, error) {
	querySq := psql.Select("t.role").From("olympics.t_role t").Where(sq.Eq{"t.name": name})
	query, args, err := querySq.ToSql()
	if err != nil {
		return 0, err
	}

	row := s.pool.QueryRowContext(ctx, query, args...)
	if err := row.Err(); err != nil {
		return 0, err
	}

	var role entities.Role
	if err := row.Scan(&role); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Role(0), nil
		}
		return 0, err
	}

	if err := row.Err(); err != nil {
		return 0, err
	}
	return role, nil
}

func (s *Storage) AddRole(ctx context.Context, name string, role entities.Role) error {
	query := `insert into olympics.t_role(name, role) values($1, $2) on conflict(name) DO UPDATE SET role = $2;`
	_, err := s.pool.Exec(query, name, role)
	return err
}
