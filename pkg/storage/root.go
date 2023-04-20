package storage

import (
	"context"
	"time"

	"olympics/pkg/core/entities"
)

type Config struct {
	DSN             string        `yaml:"dsn" config:"dsn"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	MaxConnLifetime time.Duration `yaml:"maxConnLifetime"`
	MaxConnIdleTime time.Duration `yaml:"maxConnIdleTime"`
}

type Storage interface {
	GetInfoByClient(ctx context.Context, name string) (entities.Role, error)
	AddRole(ctx context.Context, name string, role entities.Role) error
	GetAthleteInfo(ctx context.Context, name string) ([]entities.Athlete, error)
	AddAthleteEvent(ctx context.Context, athlete entities.Athlete) error
}
