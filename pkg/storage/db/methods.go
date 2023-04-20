package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"olympics/pkg/core/entities"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) GetAthleteInfo(ctx context.Context, name string) ([]entities.Athlete, error) {
	querySq := psql.Select("t.name, t.country, t.sport, t.gold, t.silver, t.bronze").
		From("olympics.t_athletes t").Where(sq.Eq{"t.name": name})
	query, args, err := querySq.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.pool.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []entities.Athlete
	for rows.Next() {
		var entry entities.Athlete
		if err := rows.Scan(&entry.Name, &entry.Sport, &entry.Country, &entry.Gold, &entry.Silver, &entry.Bronze); err != nil {
			return nil, err
		}
		result = append(result, entry)
	}
	if err := rows.Err(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (s *Storage) AddAthleteEvent(ctx context.Context, athlete entities.Athlete) error {
	query := `INSERT INTO olympics.t_athletes(name, country, sport, age, year, date, gold, silver, bronze)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	_, err := s.pool.ExecContext(ctx, query,
		athlete.Name, athlete.Country, athlete.Sport,
		athlete.Age, athlete.Year, athlete.Date,
		athlete.Gold, athlete.Silver, athlete.Bronze,
	)

	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			return fmt.Errorf("gg")
		}
		return err
	}
	return nil
}
